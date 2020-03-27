package main

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mesg-foundation/engine/ext/xos"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/protobuf/acknowledgement"
	pb "github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/runner"
	"github.com/mesg-foundation/engine/runner/builder"
	runnermodule "github.com/mesg-foundation/engine/x/runner"
	runnerrest "github.com/mesg-foundation/engine/x/runner/client/rest"
	"github.com/stretchr/testify/require"
)

var (
	testRunnerHash       hash.Hash
	testInstanceEnvHash  hash.Hash
	testRunnerAddress    sdk.AccAddress
	testServiceImageHash string
)

func testRunner(t *testing.T) {
	var (
		testInstanceEnv = xos.EnvMergeSlices(testServiceStruct.Configuration.Env, []string{"BAR=3", "REQUIRED=4"})
	)
	t.Run("hash", func(t *testing.T) {
		var res runnerrest.HashResponse
		lcdPost("runner/hash", &runnerrest.HashRequest{
			ServiceHash: testServiceHash,
			Address:     engineAddress.String(),
			Env:         testInstanceEnv,
		}, &res)
		testRunnerHash = res.RunnerHash
		testInstanceHash = res.InstanceHash
		testInstanceEnvHash = res.EnvHash
	})

	t.Run("build service image", func(t *testing.T) {
		var err error
		testServiceImageHash, err = builder.Build(cont, testServiceStruct, ipfsEndpoint)
		require.NoError(t, err)
	})

	t.Run("start", func(t *testing.T) {
		require.NoError(t, builder.Start(cont, testServiceStruct, testInstanceHash, testRunnerHash, testServiceImageHash, testInstanceEnv, engineName, enginePort))
	})

	t.Run("register", func(t *testing.T) {
		stream, err := client.EventClient.Stream(context.Background(), &pb.StreamEventRequest{
			Filter: &pb.StreamEventRequest_Filter{
				Key: "test_service_ready",
			},
		})
		require.NoError(t, err)
		acknowledgement.WaitForStreamToBeReady(stream)

		msg := runnermodule.MsgCreate{
			Owner:       engineAddress,
			ServiceHash: testServiceHash,
			EnvHash:     testInstanceEnvHash,
		}
		require.True(t, testRunnerHash.Equal(lcdBroadcastMsg(msg)))

		// wait for service to be ready
		_, err = stream.Recv()
		require.NoError(t, err)
	})

	t.Run("get", func(t *testing.T) {
		var run *runner.Runner
		lcdGet("runner/get/"+testRunnerHash.String(), &run)
		require.Equal(t, testRunnerHash, run.Hash)
		testRunnerAddress = run.Address
	})

	t.Run("list", func(t *testing.T) {
		rs := make([]*runner.Runner, 0)
		lcdGet("runner/list", &rs)
		require.Len(t, rs, 1)
		require.Equal(t, testInstanceHash, rs[0].InstanceHash)
		require.Equal(t, testRunnerHash, rs[0].Hash)
	})
}

func testDeleteRunner(t *testing.T) {
	msg := runnermodule.MsgDelete{
		Owner: engineAddress,
		Hash:  testRunnerHash,
	}
	lcdBroadcastMsg(msg)

	require.NoError(t, builder.Stop(cont, testRunnerHash, testServiceStruct.Dependencies))

	t.Run("check deletion", func(t *testing.T) {
		rs := make([]*runner.Runner, 0)
		lcdGet("runner/list", &rs)
		require.Len(t, rs, 0)
	})

	t.Run("check coins on runner", func(t *testing.T) {
		var coins sdk.Coins
		lcdGet("bank/balances/"+testRunnerAddress.String(), &coins)
		require.True(t, coins.IsZero(), coins)
	})
}
