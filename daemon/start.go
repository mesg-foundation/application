package daemon

import (
	"github.com/docker/docker/api/types/swarm"
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/docker"
	"github.com/spf13/viper"
)

// Start the docker daemon
func Start() (service *swarm.Service, err error) {
	running, err := IsRunning()
	if err != nil || running == true {
		return
	}
	network, err := docker.CreateNetwork(NamespaceNetwork(), networkDriver)
	if err != nil {
		return
	}
	return docker.StartService(&docker.ServiceOptions{
		Image:     image,
		Namespace: Namespace(),
		Ports: []docker.Port{
			docker.Port{
				Target:    50052,
				Published: 50052,
			},
		},
		Mounts: []docker.Mount{
			docker.Mount{
				Source: dockerSocket,
				Target: dockerSocket,
			},
			docker.Mount{
				Source: viper.GetString(config.MESGPath),
				Target: "/mesg",
			},
		},
		NetworksID: []string{network.ID},
	})
}
