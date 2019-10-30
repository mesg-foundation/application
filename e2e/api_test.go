package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/mesg-foundation/engine/protobuf/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var passmd = metadata.Pairs(
	"credential_username", "dev",
	"credential_passphrase", "pass",
)

type apiclient struct {
	pb.ServiceClient
	pb.EventClient
	pb.ExecutionClient
	pb.AccountClient
	pb.ProcessClient
	pb.InstanceClient
	pb.OwnershipClient
}

var client apiclient

func TestAPI(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	conn, err := grpc.DialContext(context.Background(), "localhost:50052", grpc.WithInsecure())
	require.NoError(t, err)

	client = apiclient{
		pb.NewServiceClient(conn),
		pb.NewEventClient(conn),
		pb.NewExecutionClient(conn),
		pb.NewAccountClient(conn),
		pb.NewProcessClient(conn),
		pb.NewInstanceClient(conn),
		pb.NewOwnershipClient(conn),
	}

	// wait for the first block
	time.Sleep(100 * time.Millisecond)

	// ping server to test connection
	_, err = client.ServiceClient.List(context.Background(), &pb.ListServiceRequest{})
	require.NoError(t, err)

	t.Run("account", testAccount)
	t.Run("service", testService)
	t.Run("ownership", testOwnership)
}