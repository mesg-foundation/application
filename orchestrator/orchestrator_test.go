package orchestrator

import (
	"fmt"
	"testing"

	"github.com/mesg-foundation/engine/event"
	"github.com/mesg-foundation/engine/execution"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/orchestrator/mocks"
	"github.com/mesg-foundation/engine/process"
	"github.com/mesg-foundation/engine/protobuf/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	o := New(&mocks.EventSDK{}, &mocks.ExecutionSDK{}, &mocks.ProcessSDK{})
	p := process.Process{
		Hash: hash.Int(1),
		Nodes: []*process.Process_Node{
			{Type: &process.Process_Node_Event_{Event: &process.Process_Node_Event{
				Key:          "1",
				InstanceHash: hash.Int(1),
				EventKey:     "1",
			}}},
			{Type: &process.Process_Node_Task_{Task: &process.Process_Node_Task{
				InstanceHash: hash.Int(2),
				Key:          "2",
				TaskKey:      "2",
			}}},
			{Type: &process.Process_Node_Task_{Task: &process.Process_Node_Task{
				InstanceHash: hash.Int(3),
				Key:          "3",
				TaskKey:      "3",
			}}},
		},
		Edges: []*process.Process_Edge{
			{Src: "1", Dst: "2"},
			{Src: "2", Dst: "3"},
		},
	}
	var tests = []struct {
		filter func(wf *process.Process, node *process.Process_Node) (bool, error)
		p      *process.Process
		n      *process.Process_Node
		res    bool
		err    error
	}{
		{
			filter: o.eventFilter(&event.Event{InstanceHash: hash.Int(1), Key: "1"}),
			n:      p.Nodes[0],
			res:    true,
			err:    nil,
		},
		{
			filter: o.eventFilter(&event.Event{InstanceHash: hash.Int(1), Key: "2"}),
			n:      p.Nodes[0],
			res:    false,
			err:    nil,
		},
		{
			filter: o.eventFilter(&event.Event{InstanceHash: hash.Int(2), Key: "1"}),
			n:      p.Nodes[0],
			res:    false,
			err:    nil,
		},
		{
			filter: o.eventFilter(&event.Event{InstanceHash: hash.Int(2), Key: "1"}),
			n:      p.Nodes[1],
			res:    false,
			err:    nil,
		},
		{
			filter: o.resultFilter(&execution.Execution{InstanceHash: hash.Int(1), TaskKey: "1"}),
			n: &process.Process_Node{Type: &process.Process_Node_Result_{Result: &process.Process_Node_Result{
				InstanceHash: hash.Int(1),
				TaskKey:      "1",
			}}},
			res: true,
			err: nil,
		},
		{
			filter: o.resultFilter(&execution.Execution{InstanceHash: hash.Int(1), TaskKey: "1"}),
			n: &process.Process_Node{Type: &process.Process_Node_Result_{Result: &process.Process_Node_Result{
				InstanceHash: hash.Int(1),
				TaskKey:      "2",
			}}},
			res: false,
			err: nil,
		},
		{
			filter: o.resultFilter(&execution.Execution{InstanceHash: hash.Int(1), TaskKey: "1"}),
			n: &process.Process_Node{Type: &process.Process_Node_Result_{Result: &process.Process_Node_Result{
				InstanceHash: hash.Int(2),
				TaskKey:      "1",
			}}},
			res: false,
			err: nil,
		},
		{
			filter: o.resultFilter(&execution.Execution{InstanceHash: hash.Int(1), TaskKey: "1"}),
			n:      p.Nodes[0],
			res:    false,
			err:    nil,
		},
		{
			filter: o.dependencyFilter(&execution.Execution{InstanceHash: hash.Int(3), TaskKey: "2", ProcessHash: hash.Int(1), StepID: "2"}),
			p:      &p,
			n:      p.Nodes[2],
			res:    true,
			err:    nil,
		},
		{
			filter: o.dependencyFilter(&execution.Execution{InstanceHash: hash.Int(3), TaskKey: "2", ProcessHash: hash.Int(2), StepID: "2"}),
			p:      &p,
			n:      p.Nodes[2],
			res:    false,
			err:    nil,
		},
		{
			filter: o.dependencyFilter(&execution.Execution{InstanceHash: hash.Int(3), TaskKey: "2", ProcessHash: hash.Int(1), StepID: "1"}),
			p:      &p,
			n:      p.Nodes[2],
			res:    false,
			err:    nil,
		},
		{
			filter: o.dependencyFilter(&execution.Execution{InstanceHash: hash.Int(3), TaskKey: "2", ProcessHash: hash.Int(1), StepID: "2"}),
			p:      &p,
			n:      p.Nodes[0],
			res:    false,
			err:    nil,
		},
	}
	for _, test := range tests {
		ok, err := test.filter(test.p, test.n)
		if test.err != nil {
			require.Equal(t, test.err, err)
		} else {
			require.Equal(t, ok, test.res)
		}
	}
}

func TestFindNode(t *testing.T) {
	o := New(&mocks.EventSDK{}, &mocks.ExecutionSDK{}, &mocks.ProcessSDK{})
	data := &process.Process{
		Hash: hash.Int(1),
		Nodes: []*process.Process_Node{
			{
				Type: &process.Process_Node_Event_{
					Event: &process.Process_Node_Event{
						Key: "1",
					},
				},
			},
		},
	}
	require.Len(t, o.findNodes(data, func(p *process.Process, n *process.Process_Node) (bool, error) {
		return true, nil
	}), 1)
	require.Len(t, o.findNodes(data, func(p *process.Process, n *process.Process_Node) (bool, error) {
		return n.GetEvent().Key == "1", nil
	}), 1)
	require.Len(t, o.findNodes(data, func(p *process.Process, n *process.Process_Node) (bool, error) {
		return n.GetEvent().Key == "2", nil
	}), 0)
}

// func TestProcessMap(t *testing.T) {
// 	e := &mocks.ExecutionSDK{}
// 	o := New(&mocks.EventSDK{}, e, &mocks.ProcessSDK{})
// 	exec := &execution.Execution{
// 		ProcessHash: hash.Int(1),
// 		StepID:      "1",
// 		ParentHash:  hash.Int(2),
// 		Outputs: []*types.Value{
//			&types.Value{
// 				Kind: &types.Value_StringValue{StringValue: "str"},
// 			},
// 		},
// 	}
// 	o.processMap(&process.Process_Node_Map{
// 		Outputs: []*process.Process_Node_Map_Output{},
// 	}, )
// }

func TestResolveInput(t *testing.T) {
	e := &mocks.ExecutionSDK{}
	o := New(&mocks.EventSDK{}, e, &mocks.ProcessSDK{})
	exec := &execution.Execution{
		ProcessHash: hash.Int(2),
		StepID:      "2",
		ParentHash:  hash.Int(3),
		Outputs: []*types.Value{
			{
				Kind: &types.Value_StringValue{StringValue: "str"},
			},
		},
	}
	// Different processes
	_, err := o.resolveInput(hash.Int(1), exec, "2", 0)
	require.Error(t, err)
	// Different steps, should return the value of the data
	val, err := o.resolveInput(hash.Int(2), exec, "2", 0)
	require.NoError(t, err)
	require.Equal(t, val, exec.Outputs[0])
	// Invalid execution parent hash
	e.On("Get", mock.Anything).Once().Return(nil, fmt.Errorf("err"))
	_, err = o.resolveInput(hash.Int(2), exec, "-", 0)
	require.Error(t, err)
	// Output from a previous exec
	execMock := &execution.Execution{
		StepID:      "3",
		ProcessHash: hash.Int(2),
		Outputs: []*types.Value{
			{
				Kind: &types.Value_StringValue{StringValue: "str2"},
			},
		},
	}
	e.On("Get", mock.Anything).Once().Return(execMock, nil)
	val, err = o.resolveInput(hash.Int(2), exec, "3", 0)
	require.NoError(t, err)
	require.Equal(t, val, execMock.Outputs[0])
}

func TestProcessTask(t *testing.T) {
	e := &mocks.ExecutionSDK{}
	e.On("Execute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Once().Return(nil, nil)
	o := New(&mocks.EventSDK{}, e, &mocks.ProcessSDK{})
	err := o.processTask(&process.Process_Node_Task{
		InstanceHash: hash.Int(1),
		Key:          "-",
		TaskKey:      "-",
	}, &process.Process{
		Hash: hash.Int(2),
	}, &execution.Execution{
		Hash: hash.Int(3),
	}, nil, []*types.Value{})
	require.NoError(t, err)
	e.On("Execute", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Once().Return(nil, fmt.Errorf("error"))
	err = o.processTask(&process.Process_Node_Task{
		InstanceHash: hash.Int(1),
		Key:          "-",
		TaskKey:      "-",
	}, &process.Process{
		Hash: hash.Int(2),
	}, nil, &event.Event{
		Hash: hash.Int(3),
	}, []*types.Value{})
	require.Error(t, err)
}

// func TestStart(t *testing.T) {
// 	event := &mocks.EventSDK{}
// 	exec := &mocks.ExecutionSDK{}
// 	process := &mocks.ProcessSDK{}
// 	o := New(event, exec, process)
// 	eventListener := &eventsdk.Listener{}
// 	execListener := &executionsdk.Listener{}
// 	event.On("GetStream", mock.Anything).Return(eventListener)
// 	exec.On("GetStream", mock.Anything).Return(execListener)
// 	err := o.Start()
// 	eventListener.Close()
// 	execListener.Close()
// 	require.NoError(t, err)
// }
