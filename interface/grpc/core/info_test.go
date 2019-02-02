package core

import (
	"context"
	"testing"

	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	server, closer := newServer(t)
	defer closer()

	c, err := config.Global()
	require.NoError(t, err)
	reply, err := server.Info(context.Background(), &coreapi.InfoRequest{})
	require.NoError(t, err)
	require.NotNil(t, reply)
	for i, s := range reply.Services {
		require.Equal(t, s.Sid, c.Services()[i].Sid)
	}
	require.Equal(t, reply.Address, c.Client.Address)
	require.Equal(t, reply.Image, c.Core.Image)
}
