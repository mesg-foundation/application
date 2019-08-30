package api

import (
	"context"
	"errors"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/mesg-foundation/engine/execution"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/protobuf/acknowledgement"
	"github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/protobuf/convert"
	"github.com/mesg-foundation/engine/protobuf/types"
	"github.com/mesg-foundation/engine/sdk"
	executionsdk "github.com/mesg-foundation/engine/sdk/execution"
)

// ErrNoOutput is an error when there is no output for updating execution.
var ErrNoOutput = errors.New("output not supplied")

// ExecutionServer serve execution functions.
type ExecutionServer struct {
	sdk *sdk.SDK
}

// NewExecutionServer creates a new ExecutionServer.
func NewExecutionServer(sdk *sdk.SDK) *ExecutionServer {
	return &ExecutionServer{sdk: sdk}
}

// Create creates an execution.
func (s *ExecutionServer) Create(ctx context.Context, req *api.CreateExecutionRequest) (*api.CreateExecutionResponse, error) {
	inputs := make(map[string]interface{})
	if err := convert.Marshal(req.Inputs, &inputs); err != nil {
		return nil, err
	}
	eventHash, err := hash.Random()
	if err != nil {
		return nil, err
	}
	executionHash, err := s.sdk.Execution.Execute(nil, req.InstanceHash, eventHash, nil, "", req.TaskKey, inputs, req.Tags)
	if err != nil {
		return nil, err
	}

	return &api.CreateExecutionResponse{
		Hash: executionHash,
	}, nil
}

// Get returns execution from given hash.
func (s *ExecutionServer) Get(ctx context.Context, req *api.GetExecutionRequest) (*types.Execution, error) {
	exec, err := s.sdk.Execution.Get(req.Hash)
	if err != nil {
		return nil, err
	}
	return toProtoExecution(exec)
}

// Stream returns stream of executions.
func (s *ExecutionServer) Stream(req *api.StreamExecutionRequest, resp api.Execution_StreamServer) error {
	var f *executionsdk.Filter

	if req.Filter != nil {
		var statuses []execution.Status
		for _, status := range req.Filter.Statuses {
			statuses = append(statuses, execution.Status(status))
		}

		f = &executionsdk.Filter{
			InstanceHash: req.Filter.InstanceHash,
			Statuses:     statuses,
			Tags:         req.Filter.Tags,
			TaskKey:      req.Filter.TaskKey,
		}
	}

	stream := s.sdk.Execution.GetStream(f)
	defer stream.Close()

	// send header to notify client that the stream is ready.
	if err := acknowledgement.SetStreamReady(resp); err != nil {
		return err
	}

	for exec := range stream.C {
		e, err := toProtoExecution(exec)
		if err != nil {
			return err
		}

		if err := resp.Send(e); err != nil {
			return err
		}
	}

	return nil
}

// Update updates execution from given hash.
func (s *ExecutionServer) Update(ctx context.Context, req *api.UpdateExecutionRequest) (*api.UpdateExecutionResponse, error) {
	var err error
	switch res := req.Result.(type) {
	case *api.UpdateExecutionRequest_Outputs:
		outputs := make(map[string]interface{})
		if err := convert.Marshal(res.Outputs, &outputs); err != nil {
			return nil, err
		}

		err = s.sdk.Execution.Update(req.Hash, outputs, nil)
	case *api.UpdateExecutionRequest_Error:
		err = s.sdk.Execution.Update(req.Hash, nil, errors.New(res.Error))
	default:
		err = ErrNoOutput
	}

	if err != nil {
		return nil, err
	}
	return &api.UpdateExecutionResponse{}, nil

}

func toProtoExecution(exec *execution.Execution) (*types.Execution, error) {
	inputs := &structpb.Struct{}
	if err := convert.Unmarshal(exec.Inputs, inputs); err != nil {
		return nil, err
	}

	outputs := &structpb.Struct{}
	if err := convert.Unmarshal(exec.Outputs, outputs); err != nil {
		return nil, err
	}

	return &types.Execution{
		Hash:         exec.Hash,
		ProcessHash: exec.ProcessHash,
		ParentHash:   exec.ParentHash,
		EventHash:    exec.EventHash,
		Status:       types.Status(exec.Status),
		InstanceHash: exec.InstanceHash,
		TaskKey:      exec.TaskKey,
		Inputs:       inputs,
		Outputs:      outputs,
		Tags:         exec.Tags,
		Error:        exec.Error,
		StepID:       exec.StepID,
	}, nil
}
