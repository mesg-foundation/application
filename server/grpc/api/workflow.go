package api

import (
	"context"
	"fmt"

	"github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/protobuf/types"
	"github.com/mesg-foundation/engine/sdk"
	"github.com/mesg-foundation/engine/workflow"
)

// WorkflowServer is the type to aggregate all Service APIs.
type WorkflowServer struct {
	sdk *sdk.SDK
}

// NewWorkflowServer creates a new WorkflowServer.
func NewWorkflowServer(sdk *sdk.SDK) *WorkflowServer {
	return &WorkflowServer{sdk: sdk}
}

// Create creates a new service from definition.
func (s *WorkflowServer) Create(ctx context.Context, req *api.CreateWorkflowRequest) (*api.CreateWorkflowResponse, error) {
	wf, err := fromProtoWorkflow(&types.Workflow{
		Key:     req.Key,
		Trigger: req.Trigger,
		Nodes:   req.Nodes,
		Edges:   req.Edges,
	})
	if err != nil {
		return nil, err
	}
	wf, err = s.sdk.Workflow.Create(wf)
	if err != nil {
		return nil, err
	}
	return &api.CreateWorkflowResponse{Hash: wf.Hash}, nil
}

// Delete deletes service by hash or sid.
func (s *WorkflowServer) Delete(ctx context.Context, request *api.DeleteWorkflowRequest) (*api.DeleteWorkflowResponse, error) {
	return &api.DeleteWorkflowResponse{}, s.sdk.Workflow.Delete(request.Hash)
}

// Get returns service from given hash.
func (s *WorkflowServer) Get(ctx context.Context, req *api.GetWorkflowRequest) (*types.Workflow, error) {
	wf, err := s.sdk.Workflow.Get(req.Hash)
	if err != nil {
		return nil, err
	}
	return toProtoWorkflow(wf), nil
}

// List returns all workflows.
func (s *WorkflowServer) List(ctx context.Context, req *api.ListWorkflowRequest) (*api.ListWorkflowResponse, error) {
	workflows, err := s.sdk.Workflow.List()
	if err != nil {
		return nil, err
	}
	wfs := toProtoWorkflows(workflows)
	return &api.ListWorkflowResponse{
		Workflows: wfs,
	}, nil
}

func fromProtoFilters(filters []*types.Workflow_Trigger_Filter) []*workflow.TriggerFilter {
	fs := make([]*workflow.TriggerFilter, len(filters))
	for i, filter := range filters {
		var predicate workflow.Predicate
		// switch filter.Predicate {
		if filter.Predicate == types.Workflow_Trigger_Filter_EQ {
			predicate = workflow.EQ
		}
		fs[i] = &workflow.TriggerFilter{
			Key:       filter.Key,
			Predicate: predicate,
			Value:     filter.Value,
		}
	}
	return fs
}

func fromProtoWorkflowNodes(nodes []*types.Workflow_Node) []workflow.Node {
	res := make([]workflow.Node, len(nodes))
	for i, node := range nodes {
		res[i] = workflow.Node{
			Key:          node.Key,
			InstanceHash: node.InstanceHash,
			TaskKey:      node.TaskKey,
		}
	}
	return res
}

func fromProtoWorkflowEdges(edges []*types.Workflow_Edge) ([]workflow.Edge, error) {
	res := make([]workflow.Edge, len(edges))
	for i, edge := range edges {
		inputs := make([]*workflow.Input, len(edge.Inputs))
		for j, input := range edge.Inputs {
			in := &workflow.Input{Key: input.Key}
			switch x := input.Value.(type) {
			case *types.Workflow_Edge_Input_Ref:
				in.Ref = &workflow.InputReference{
					NodeKey: input.GetRef().NodeKey,
					Key:     input.GetRef().Key,
				}
			default:
				return nil, fmt.Errorf("input has unexpected type %T", x)
			}
			inputs[j] = in
		}
		res[i] = workflow.Edge{
			Src:    edge.Src,
			Dst:    edge.Dst,
			Inputs: inputs,
		}
	}
	return res, nil
}

func fromProtoWorkflow(wf *types.Workflow) (*workflow.Workflow, error) {
	nodes := fromProtoWorkflowNodes(wf.Nodes)
	edges, err := fromProtoWorkflowEdges(wf.Edges)
	if err != nil {
		return nil, err
	}
	trigger := workflow.Trigger{
		InstanceHash: wf.Trigger.InstanceHash,
		NodeKey:      wf.Trigger.NodeKey,
		Filters:      fromProtoFilters(wf.Trigger.Filters),
	}
	switch x := wf.Trigger.Key.(type) {
	case *types.Workflow_Trigger_EventKey:
		trigger.EventKey = x.EventKey
	case *types.Workflow_Trigger_TaskKey:
		trigger.TaskKey = x.TaskKey
	default:
		return nil, fmt.Errorf("workflow trigger key has unexpected type %T", x)
	}
	return &workflow.Workflow{
		Key:     wf.Key,
		Trigger: trigger,
		Nodes:   nodes,
		Edges:   edges,
	}, nil
}

func toProtoFilters(filters []*workflow.TriggerFilter) []*types.Workflow_Trigger_Filter {
	fs := make([]*types.Workflow_Trigger_Filter, len(filters))
	for i, filter := range filters {
		var predicate types.Workflow_Trigger_Filter_Predicate
		// switch filter.Predicate {
		if filter.Predicate == workflow.EQ {
			predicate = types.Workflow_Trigger_Filter_EQ
		}
		fs[i] = &types.Workflow_Trigger_Filter{
			Key:       filter.Key,
			Predicate: predicate,
			Value:     filter.Value.(string),
		}
	}
	return fs
}

func toProtoWorkflowNodes(nodes []workflow.Node) []*types.Workflow_Node {
	res := make([]*types.Workflow_Node, len(nodes))
	for i, node := range nodes {
		res[i] = &types.Workflow_Node{
			Key:          node.Key,
			InstanceHash: node.InstanceHash,
			TaskKey:      node.TaskKey,
		}
	}
	return res
}

func toProtoWorkflowEdges(edges []workflow.Edge) []*types.Workflow_Edge {
	res := make([]*types.Workflow_Edge, len(edges))
	for i, edge := range edges {
		inputs := make([]*types.Workflow_Edge_Input, len(edge.Inputs))
		for j, input := range edge.Inputs {
			in := &types.Workflow_Edge_Input{Key: input.Key}
			if input.Ref != nil {
				in.Value = &types.Workflow_Edge_Input_Ref{
					Ref: &types.Workflow_Edge_Input_Reference{
						NodeKey: input.Ref.NodeKey,
						Key:     input.Ref.Key,
					},
				}
			}
			inputs[j] = in
		}
		res[i] = &types.Workflow_Edge{
			Src:    edge.Src,
			Dst:    edge.Dst,
			Inputs: inputs,
		}
	}
	return res
}

func toProtoWorkflow(wf *workflow.Workflow) *types.Workflow {
	trigger := &types.Workflow_Trigger{
		InstanceHash: wf.Trigger.InstanceHash,
		Filters:      toProtoFilters(wf.Trigger.Filters),
		NodeKey:      wf.Trigger.NodeKey,
	}
	if wf.Trigger.TaskKey != "" {
		trigger.Key = &types.Workflow_Trigger_TaskKey{TaskKey: wf.Trigger.TaskKey}
	}
	if wf.Trigger.EventKey != "" {
		trigger.Key = &types.Workflow_Trigger_EventKey{EventKey: wf.Trigger.EventKey}
	}
	return &types.Workflow{
		Hash:    wf.Hash,
		Key:     wf.Key,
		Trigger: trigger,
		Nodes:   toProtoWorkflowNodes(wf.Nodes),
		Edges:   toProtoWorkflowEdges(wf.Edges),
	}
}

func toProtoWorkflows(workflows []*workflow.Workflow) []*types.Workflow {
	wfs := make([]*types.Workflow, len(workflows))
	for i, wf := range workflows {
		wfs[i] = toProtoWorkflow(wf)
	}
	return wfs
}
