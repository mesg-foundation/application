package core

import (
	"errors"
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
		statuses = make(chan api.DeployStatus)
		wg       sync.WaitGroup

		service         *service.Service
		validationError *importer.ValidationError
		err             error
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		sendDeployStatus(statuses, stream)
	}()

	// read first requesest from stream and check if it's url or tarball
	in, err := stream.Recv()
	if err != nil {
		return err
	}

	// env must go always with first package
	env := in.GetEnv()

	if url := in.GetUrl(); url != "" {
		service, validationError, err = s.api.DeployServiceFromURL(url, env, api.DeployServiceStatusOption(statuses))
	} else {
		tarball := &deployChunkReader{
			stream: stream,
			buf:    in.GetChunk(),
		}
		service, validationError, err = s.api.DeployService(tarball, env, api.DeployServiceStatusOption(statuses))
	}
	wg.Wait()

	if err != nil {
		return err
	}
	if validationError != nil {
		return stream.Send(&coreapi.DeployServiceReply{
			Value: &coreapi.DeployServiceReply_ValidationError{
				ValidationError: validationError.Error(),
			},
		})
	}

	return stream.Send(&coreapi.DeployServiceReply{
		Value: &coreapi.DeployServiceReply_ServiceID{
			ServiceID: service.Hash,
		},
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

type deployChunkReader struct {
	stream coreapi.Core_DeployServiceServer

	buf []byte
	i   int
}

func (r *deployChunkReader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.buf) {
		in, err := r.stream.Recv()
		if err != nil {
			return 0, err
		}
		if in.GetUrl() != "" {
			return 0, errors.New("deploy: got url after tarball stream")
		}
		r.buf = in.GetChunk()
		r.i = 0
	}
	n = copy(p, r.buf[r.i:])
	r.i += n
	return n, nil
}
