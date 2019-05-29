package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/execution"
	"github.com/mesg-foundation/core/protobuf/acknowledgement"
	"github.com/mesg-foundation/core/protobuf/serviceapi"
)

var inProgressFilter = &api.ExecutionFilter{
	Statuses: []execution.Status{execution.InProgress},
}

// Server binds all api functions.
type Server struct {
	api *api.API
}

// NewServer creates a new Server.
func NewServer(api *api.API) *Server {
	return &Server{api: api}
}

// EmitEvent permits to send and event to anyone who subscribed to it.
func (s *Server) EmitEvent(context context.Context, request *serviceapi.EmitEventRequest) (*serviceapi.EmitEventReply, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(request.EventData), &data); err != nil {
		return nil, err
	}
	return &serviceapi.EmitEventReply{}, s.api.EmitEvent(request.Token, request.EventKey, data)
}

// ListenTask creates a stream that will send data for every task to execute.
func (s *Server) ListenTask(request *serviceapi.ListenTaskRequest, stream serviceapi.Service_ListenTaskServer) error {
	ln, err := s.api.ListenExecution(request.Token, inProgressFilter)
	if err != nil {
		return err
	}
	defer ln.Close()

	// send header to notify client that the stream is ready.
	if err := acknowledgement.SetStreamReady(stream); err != nil {
		return err
	}

	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case execution := <-ln.C:
			inputs, err := json.Marshal(execution.Inputs)
			if err != nil {
				return err
			}

			if err := stream.Send(&serviceapi.TaskData{
				ExecutionHash: string(execution.Hash),
				TaskKey:       execution.TaskKey,
				InputData:     string(inputs),
			}); err != nil {
				return err
			}
		}
	}
}

// SubmitResult submits results of an execution.
func (s *Server) SubmitResult(context context.Context, request *serviceapi.SubmitResultRequest) (*serviceapi.SubmitResultReply, error) {
	switch res := request.Result.(type) {
	case *serviceapi.SubmitResultRequest_OutputData:
		return &serviceapi.SubmitResultReply{}, s.api.SubmitResult([]byte(request.ExecutionHash), []byte(res.OutputData), nil)
	case *serviceapi.SubmitResultRequest_Error:
		return &serviceapi.SubmitResultReply{}, s.api.SubmitResult([]byte(request.ExecutionHash), nil, errors.New(res.Error))
	}
	return &serviceapi.SubmitResultReply{}, nil
}
