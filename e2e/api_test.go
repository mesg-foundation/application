package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/mesg-foundation/engine/app"
	"github.com/mesg-foundation/engine/config"
	"github.com/mesg-foundation/engine/cosmos"
	pb "github.com/mesg-foundation/engine/protobuf/api"
	"github.com/stretchr/testify/require"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"google.golang.org/grpc"
)

type apiclient struct {
	pb.ServiceClient
	pb.EventClient
	pb.ExecutionClient
	pb.ProcessClient
	pb.InstanceClient
	pb.OwnershipClient
	pb.RunnerClient
}

var (
	minExecutionPrice     sdk.Coins
	client                apiclient
	cclient               *cosmos.Client
	cdc                   = app.MakeCodec()
	processInitialBalance = sdk.NewCoins(sdk.NewInt64Coin("atto", 10000000))
	engineAddress         sdk.AccAddress
)

const (
	lcdEndpoint        = "http://127.0.0.1:1317/"
	lcdPostContentType = "application/json"
)

func lcdGet(t *testing.T, path string, ptr interface{}) {
	resp, err := http.Get(lcdEndpoint + path)
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	cosResp := rest.ResponseWithHeight{}
	require.NoError(t, cdc.UnmarshalJSON(body, &cosResp))
	if len(cosResp.Result) > 0 {
		require.NoError(t, cdc.UnmarshalJSON(cosResp.Result, ptr))
	}
}

func lcdPost(t *testing.T, path string, req interface{}, ptr interface{}) {
	reqBody, err := cdc.MarshalJSON(req)
	require.NoError(t, err)
	resp, err := http.Post(lcdEndpoint+path, lcdPostContentType, bytes.NewReader(reqBody))
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	cosResp := rest.ResponseWithHeight{}
	require.NoError(t, cdc.UnmarshalJSON(body, &cosResp))
	if len(cosResp.Result) > 0 {
		require.NoError(t, cdc.UnmarshalJSON(cosResp.Result, ptr))
	}
}

func TestAPI(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cfg, err := config.New()
	require.NoError(t, err)

	minExecutionPrice, err = sdk.ParseCoins(cfg.DefaultExecutionPrice)
	require.NoError(t, err)

	cosmos.CustomizeConfig(cfg)

	conn, err := grpc.DialContext(context.Background(), "localhost:50052", grpc.WithInsecure())
	require.NoError(t, err)

	// change and recreate cosmos relative path because CI dir permissions
	cfg.Cosmos.RelativePath = "e2e.cosmos"
	err = os.MkdirAll(filepath.Join(cfg.Path, cfg.Cosmos.RelativePath), os.FileMode(0755))
	require.NoError(t, err)

	kb, err := cosmos.NewKeybase(filepath.Join(cfg.Path, cfg.Cosmos.RelativePath))
	require.NoError(t, err)
	if cfg.Account.Mnemonic != "" {
		acc, err := kb.CreateAccount(cfg.Account.Name, cfg.Account.Mnemonic, "", cfg.Account.Password, keys.CreateHDPath(cfg.Account.Number, cfg.Account.Index).String(), cosmos.DefaultAlgo)
		require.NoError(t, err)
		engineAddress = acc.GetAddress()
	}

	httpclient, err := rpcclient.NewHTTP("http://localhost:26657", "/websocket")
	require.NoError(t, err)
	require.NoError(t, httpclient.Start())
	defer httpclient.Stop()
	cclient, err = cosmos.NewClient(httpclient, cdc, kb, cfg.DevGenesis.ChainID, cfg.Account.Name, cfg.Account.Password, cfg.Cosmos.MinGasPrices)
	require.NoError(t, err)

	client = apiclient{
		pb.NewServiceClient(conn),
		pb.NewEventClient(conn),
		pb.NewExecutionClient(conn),
		pb.NewProcessClient(conn),
		pb.NewInstanceClient(conn),
		pb.NewOwnershipClient(conn),
		pb.NewRunnerClient(conn),
	}

	// ping server to test connection
	_, err = client.ServiceClient.List(context.Background(), &pb.ListServiceRequest{})
	require.NoError(t, err)

	// basic tests
	t.Run("service", testService)
	t.Run("runner", testRunner)
	t.Run("process", testProcess)
	t.Run("instance", testInstance)
	t.Run("event", testEvent)
	t.Run("execution", testExecution)
	t.Run("orchestrator", testOrchestrator)
	t.Run("runner/delete", testDeleteRunner)
	t.Run("complex-service", testComplexService)
}
