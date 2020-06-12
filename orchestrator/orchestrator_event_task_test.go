package orchestrator

import (
	"context"
	"testing"
	"time"

	"github.com/mesg-foundation/engine/execution"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/process"
	"github.com/mesg-foundation/engine/protobuf/types"
	"github.com/stretchr/testify/require"
)

func TestOrchestratorEventTask(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	orch, ep, store, _, testInstanceHash, _, _, execChan := newTestOrchestrator(ctx, t)
	defer orch.Stop()

	var (
		testProcessHash hash.Hash
		err             error
	)
	t.Run("create process", func(t *testing.T) {
		testProcessHash, err = store.CreateProcess(
			context.Background(),
			"result-task-process",
			[]*process.Process_Node{
				{
					Key: "n0",
					Type: &process.Process_Node_Event_{
						Event: &process.Process_Node_Event{
							InstanceHash: testInstanceHash,
							EventKey:     "event_trigger",
						},
					},
				},
				{
					Key: "n1",
					Type: &process.Process_Node_Task_{
						Task: &process.Process_Node_Task{
							InstanceHash: testInstanceHash,
							TaskKey:      "task1",
						},
					},
				},
			}, []*process.Process_Edge{
				{Src: "n0", Dst: "n1"},
			},
		)
		require.NoError(t, err)
	})
	t.Run("trigger process", func(t *testing.T) {
		_, err := ep.Publish(
			testInstanceHash,
			"event_trigger",
			&types.Struct{
				Fields: map[string]*types.Value{
					"msg": {
						Kind: &types.Value_StringValue{
							StringValue: "foo_1",
						},
					},
					"timestamp": {
						Kind: &types.Value_NumberValue{
							NumberValue: float64(time.Now().Unix()),
						},
					},
				},
			},
		)
		require.NoError(t, err)
	})
	t.Run("check created execution", func(t *testing.T) {
		exec := <-execChan
		require.NoError(t, err)
		require.Equal(t, "task1", exec.TaskKey)
		require.Equal(t, "n1", exec.NodeKey)
		require.True(t, testProcessHash.Equal(exec.ProcessHash))
		require.Equal(t, execution.Status_InProgress, exec.Status)
		require.Equal(t, "foo_1", exec.Inputs.Fields["msg"].GetStringValue())
	})
}
