package intergration_test

import (
	"context"
	"testing"

	"github.com/mesg-foundation/core/daemon"
	"github.com/mesg-foundation/core/database/services"
	"github.com/stvp/assert"

	"github.com/mesg-foundation/core/api/core"
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func TestSharedDatabse(t *testing.T) {
	daemon.Start()
	defer daemon.Stop()
	err := <-daemon.WaitForContainerToRun()
	assert.Nil(t, err)
	connection, err := grpc.Dial(viper.GetString(config.APIClientTarget), grpc.WithInsecure())
	assert.Nil(t, err)
	cli := core.NewCoreClient(connection)
	reply, err := cli.DeployService(context.Background(), &core.DeployServiceRequest{
		Service: &service.Service{
			Name: "TestSharedDatabse",
			Dependencies: map[string]*service.Dependency{
				"test": &service.Dependency{
					Image: "nginx",
				},
			},
		},
	})
	assert.Nil(t, err)
	service, err := services.Get(reply.ServiceID)
	defer services.Delete(reply.ServiceID)
	assert.Nil(t, err)
	assert.Equal(t, service.Name, "TestSharedDatabse")
}
