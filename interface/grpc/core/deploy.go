package core

import (
	"sync"

	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/protobuf/coreapi"
	service "github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/service/importer"
)

// DeployService deploys a service from Git URL or service.tar.gz file. It'll send status
// events during the process and finish with sending service id or validation error.
func (s *Server) DeployService(stream coreapi.Core_DeployServiceServer) error {
	var (
		sr         = newDeployServiceStreamReader(stream)
		statuses   = make(chan api.DeployStatus)
		wgStatuses sync.WaitGroup
	)

	wgStatuses.Add(1)
	go func() {
		defer wgStatuses.Done()
		sendDeployStatus(statuses, stream)
	}()

	deployOptions := []api.DeployServiceOption{
		api.DeployServiceStatusOption(statuses),
	}

	// receive the 1th message in the stream.
	// it can be the deployment options or service's url/first chunk.
	if err := sr.RecvMessage(); err != nil {
		return err
	}
	var env map[string]string
	if sr.Options != nil {
		env = sr.Options.Env
	}

	var (
		service         *service.Service
		validationError *importer.ValidationError
		err             error
	)
	// receive the 2nd message in the stream if first message was the options.
	// message can contain url or first chunk of the service.
	if sr.Options != nil {
		if err := sr.RecvMessage(); err != nil {
			return err
		}
	}
	if sr.URL != "" {
		service, validationError, err = s.api.DeployServiceFromURL(sr.URL, env, deployOptions...)
	} else {
		service, validationError, err = s.api.DeployService(sr, env, deployOptions...)
	}

	// wait for statuses to be sent first otherwise sending multiple messages at the
	// same time may cause messages to be sent in different order.
	wgStatuses.Wait()

	if err != nil {
		return err
	}

	if validationError != nil {
		return stream.Send(&coreapi.DeployServiceReply{
			Value: &coreapi.DeployServiceReply_ValidationError{ValidationError: validationError.Error()},
		})
	}

	return stream.Send(&coreapi.DeployServiceReply{
		Value: &coreapi.DeployServiceReply_ServiceID{ServiceID: service.Hash},
	})
}

func sendDeployStatus(statuses chan api.DeployStatus, stream coreapi.Core_DeployServiceServer) {
	for status := range statuses {
		var typ coreapi.DeployServiceReply_Status_Type
		switch status.Type {
		case api.Running:
			typ = coreapi.DeployServiceReply_Status_RUNNING
		case api.DonePositive:
			typ = coreapi.DeployServiceReply_Status_DONE_POSITIVE
		case api.DoneNegative:
			typ = coreapi.DeployServiceReply_Status_DONE_NEGATIVE
		}
		stream.Send(&coreapi.DeployServiceReply{
			Value: &coreapi.DeployServiceReply_Status_{
				Status: &coreapi.DeployServiceReply_Status{
					Message: status.Message,
					Type:    typ,
				},
			},
		})
	}
}

type deployServiceStreamReader struct {
	stream coreapi.Core_DeployServiceServer

	// Options is the deployment options.
	Options *coreapi.DeployServiceRequest_Options

	// URL of the service.
	URL string

	// chunk of the service.
	chunk []byte
	i     int64
}

func newDeployServiceStreamReader(stream coreapi.Core_DeployServiceServer) *deployServiceStreamReader {
	return &deployServiceStreamReader{
		stream: stream,
	}
}

// RecvMessage receives the next message in gRPC stream.
func (r *deployServiceStreamReader) RecvMessage() error {
	message, err := r.stream.Recv()
	if err != nil {
		return err
	}
	r.Options = message.GetOptions()
	r.URL = message.GetUrl()
	r.chunk = message.GetChunk()
	return nil
}

// Read reads service chunks to deploy.
func (r *deployServiceStreamReader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.chunk)) {
		if err := r.RecvMessage(); err != nil {
			return 0, err
		}
		r.i = 0
		return r.Read(p)
	}
	n = copy(p, r.chunk[r.i:])
	r.i += int64(n)
	return n, nil
}
