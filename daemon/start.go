package daemon

import (
	"path/filepath"

	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/container"
	"github.com/spf13/viper"
)

// Start the docker core
func Start() (serviceID string, err error) {
	running, err := IsRunning()
	if err != nil || running == true {
		return
	}
	spec, err := serviceSpec()
	if err != nil {
		return
	}
	serviceID, err = container.StartService(spec)
	return
}

func serviceSpec() (spec container.ServiceOptions, err error) {
	sharedNetworkID, err := container.SharedNetworkID()
	if err != nil {
		return
	}
	spec = container.ServiceOptions{
		Namespace: Namespace(),
		Image:     viper.GetString(config.CoreImage),
		Env: []string{
			"MESG.PATH=/mesg",
			"API.SERVICE.SOCKETPATH=" + filepath.Join(viper.GetString(config.MESGPath), "server.sock"),
			"SERVICE.PATH.HOST=" + filepath.Join(viper.GetString(config.MESGPath), "services"),
		},
		Mounts: []container.Mount{
			container.Mount{
				Source: dockerSocket,
				Target: dockerSocket,
			},
			container.Mount{
				Source: viper.GetString(config.MESGPath),
				Target: "/mesg",
			},
		},
		Ports: []container.Port{
			container.Port{
				Target:    50052,
				Published: 50052,
			},
		},
		NetworksID: []string{sharedNetworkID},
	}
	return
}
