package grpc

import (
	"net"
	"sync"

	"github.com/mesg-foundation/core/systemservices"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/interface/grpc/core"
	"github.com/mesg-foundation/core/interface/grpc/service"
	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/mesg-foundation/core/protobuf/serviceapi"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server contains the server config.
type Server struct {
	api *api.API
	ss  *systemservices.SystemServices

	instance *grpc.Server
	closed   bool
	mi       sync.Mutex // protects startup.

	network string
	address string
}

// New returns a new gRPC server.
func New(address string, api *api.API, ss *systemservices.SystemServices) *Server {
	return &Server{
		api:     api,
		ss:      ss,
		address: address,
		network: "tcp",
	}
}

// listen listens for connections.
func (s *Server) listen() (net.Listener, error) {
	s.mi.Lock()
	defer s.mi.Unlock()

	if s.closed {
		return nil, &alreadyClosedError{}
	}

	ln, err := net.Listen(s.network, s.address)
	if err != nil {
		return nil, err
	}

	s.instance = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(logrus.NewEntry(logrus.StandardLogger())),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logrus.StandardLogger())),
		)),
	)
	if err := s.register(); err != nil {
		return nil, err
	}

	logrus.Info("server listens on ", ln.Addr())
	return ln, nil
}

// Serve starts the server and listens for client connections.
func (s *Server) Serve() error {
	ln, err := s.listen()
	if err != nil {
		return err
	}

	// TODO: check if server still on after a connection throw an error. otherwise, add a for around serve
	return s.instance.Serve(ln)
}

// Close gracefully closes the server.
func (s *Server) Close() {
	s.mi.Lock()
	defer s.mi.Unlock()
	if s.closed {
		return
	}
	if s.instance != nil {
		s.instance.GracefulStop()
	}
	s.closed = true
}

// register all server
func (s *Server) register() error {
	coreServer := core.NewServer(s.api, s.ss)
	serviceServer := service.NewServer(s.api)

	serviceapi.RegisterServiceServer(s.instance, serviceServer)
	coreapi.RegisterCoreServer(s.instance, coreServer)

	reflection.Register(s.instance)
	return nil
}

type alreadyClosedError struct{}

func (e *alreadyClosedError) Error() string {
	return "already closed"
}
