package main

import (
	"context"
	"testing"
	"time"

	"github.com/mesg-foundation/engine/execution"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/process"
	"github.com/mesg-foundation/engine/protobuf/acknowledgement"
	pb "github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/protobuf/types"
	processmodule "github.com/mesg-foundation/engine/x/process"
	"github.com/stretchr/testify/require"
)

func testOrchestratorFilterPathNested(instanceHash hash.Hash) func(t *testing.T) {
	return func(t *testing.T) {
		var (
			processHash     hash.Hash
			executionStream pb.Execution_StreamClient
		)

		t.Run("create process", func(t *testing.T) {
			processHash = lcdBroadcastMsg(t, processmodule.MsgCreate{
				Owner: engineAddress,
				Name:  "filter",
				Nodes: []*process.Process_Node{
					{
						Key: "n0",
						Type: &process.Process_Node_Event_{
							Event: &process.Process_Node_Event{
								InstanceHash: instanceHash,
								EventKey:     "test_event_complex",
							},
						},
					},
					{
						Key: "n1",
						Type: &process.Process_Node_Filter_{
							Filter: &process.Process_Node_Filter{
								Conditions: []process.Process_Node_Filter_Condition{
									{
										Ref: &process.Process_Node_Reference{
											NodeKey: "n0",
											Path: &process.Process_Node_Reference_Path{
												Selector: &process.Process_Node_Reference_Path_Key{
													Key: "msg",
												},
												Path: &process.Process_Node_Reference_Path{
													Selector: &process.Process_Node_Reference_Path_Key{
														Key: "msg",
													},
												},
											},
										},
										Predicate: process.Process_Node_Filter_Condition_EQ,
										Value: &types.Value{
											Kind: &types.Value_StringValue{
												StringValue: "shouldMatch",
											},
										},
									},
									{
										Ref: &process.Process_Node_Reference{
											NodeKey: "n0",
											Path: &process.Process_Node_Reference_Path{
												Selector: &process.Process_Node_Reference_Path_Key{
													Key: "msg",
												},
												Path: &process.Process_Node_Reference_Path{
													Selector: &process.Process_Node_Reference_Path_Key{
														Key: "timestamp",
													},
												},
											},
										},
										Predicate: process.Process_Node_Filter_Condition_GT,
										Value: &types.Value{
											Kind: &types.Value_NumberValue{
												NumberValue: 10,
											},
										},
									},
									{
										Ref: &process.Process_Node_Reference{
											NodeKey: "n0",
											Path: &process.Process_Node_Reference_Path{
												Selector: &process.Process_Node_Reference_Path_Key{
													Key: "msg",
												},
												Path: &process.Process_Node_Reference_Path{
													Selector: &process.Process_Node_Reference_Path_Key{
														Key: "array",
													},
												},
											},
										},
										Predicate: process.Process_Node_Filter_Condition_EQ,
										Value: &types.Value{
											Kind: &types.Value_ListValue{
												ListValue: &types.ListValue{
													Values: []*types.Value{
														{
															Kind: &types.Value_StringValue{
																StringValue: "one",
															},
														},
														{
															Kind: &types.Value_StringValue{
																StringValue: "two",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Key: "n2",
						Type: &process.Process_Node_Task_{
							Task: &process.Process_Node_Task{
								InstanceHash: instanceHash,
								TaskKey:      "task_complex",
							},
						},
					},
				},
				Edges: []*process.Process_Edge{
					{Src: "n0", Dst: "n1"},
					{Src: "n1", Dst: "n2"},
				},
			})
		})
		t.Run("create execution stream", func(t *testing.T) {
			var err error
			executionStream, err = client.ExecutionClient.Stream(context.Background(), &pb.StreamExecutionRequest{})
			require.NoError(t, err)
			acknowledgement.WaitForStreamToBeReady(executionStream)
		})
		t.Run("pass filter", func(t *testing.T) {
			t.Run("trigger process", func(t *testing.T) {
				_, err := client.EventClient.Create(context.Background(), &pb.CreateEventRequest{
					InstanceHash: instanceHash,
					Key:          "test_event_complex",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"msg": {
								Kind: &types.Value_StructValue{
									StructValue: &types.Struct{
										Fields: map[string]*types.Value{
											"msg": {
												Kind: &types.Value_StringValue{
													StringValue: "shouldMatch",
												},
											},
											"timestamp": {
												Kind: &types.Value_NumberValue{
													NumberValue: float64(time.Now().Unix()),
												},
											},
											"array": {
												Kind: &types.Value_ListValue{
													ListValue: &types.ListValue{
														Values: []*types.Value{
															{
																Kind: &types.Value_StringValue{
																	StringValue: "one",
																},
															},
															{
																Kind: &types.Value_StringValue{
																	StringValue: "two",
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				})
				require.NoError(t, err)
			})
			t.Run("check in progress execution", func(t *testing.T) {
				exec, err := executionStream.Recv()
				require.NoError(t, err)
				require.Equal(t, "task_complex", exec.TaskKey)
				require.Equal(t, "n2", exec.NodeKey)
				require.True(t, processHash.Equal(exec.ProcessHash))
				require.Equal(t, execution.Status_InProgress, exec.Status)
				require.Equal(t, "shouldMatch", exec.Inputs.Fields["msg"].GetStructValue().Fields["msg"].GetStringValue())
			})
			t.Run("check completed execution", func(t *testing.T) {
				exec, err := executionStream.Recv()
				require.NoError(t, err)
				require.Equal(t, "task_complex", exec.TaskKey)
				require.Equal(t, "n2", exec.NodeKey)
				require.True(t, processHash.Equal(exec.ProcessHash))
				require.Equal(t, execution.Status_Completed, exec.Status)
				require.Equal(t, "shouldMatch", exec.Outputs.Fields["msg"].GetStructValue().Fields["msg"].GetStringValue())
				require.NotEmpty(t, exec.Outputs.Fields["msg"].GetStructValue().Fields["timestamp"].GetNumberValue())
			})
		})
		t.Run("stop at filter", func(t *testing.T) {
			t.Run("trigger process", func(t *testing.T) {
				_, err := client.EventClient.Create(context.Background(), &pb.CreateEventRequest{
					InstanceHash: instanceHash,
					Key:          "test_event_complex",
					Data: &types.Struct{
						Fields: map[string]*types.Value{
							"msg": {
								Kind: &types.Value_StructValue{
									StructValue: &types.Struct{
										Fields: map[string]*types.Value{
											"msg": {
												Kind: &types.Value_StringValue{
													StringValue: "shouldNotMatch",
												},
											},
											"timestamp": {
												Kind: &types.Value_NumberValue{
													NumberValue: float64(time.Now().Unix()),
												},
											},
											"array": {
												Kind: &types.Value_ListValue{
													ListValue: &types.ListValue{
														Values: []*types.Value{
															{
																Kind: &types.Value_StringValue{
																	StringValue: "one",
																},
															},
															{
																Kind: &types.Value_StringValue{
																	StringValue: "two",
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				})
				require.NoError(t, err)
			})
			t.Run("wait 5 sec to check execution is not created", func(t *testing.T) {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
				defer cancel()
				executionStream, err := client.ExecutionClient.Stream(ctx, &pb.StreamExecutionRequest{})
				require.NoError(t, err)
				_, err = executionStream.Recv()
				require.Contains(t, err.Error(), context.DeadlineExceeded.Error())
			})
		})
		t.Run("delete process", func(t *testing.T) {
			lcdBroadcastMsg(t, processmodule.MsgDelete{
				Owner: engineAddress,
				Hash:  processHash,
			})
		})
	}
}
