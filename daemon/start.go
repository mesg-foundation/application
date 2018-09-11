package daemon

import (
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/x/xnet"
	"github.com/spf13/viper"
)

// Start starts the docker core.
func Start() (serviceID string, err error) {
	status, err := Status()
	if err != nil || status == container.RUNNING {
		return "", err
	}
	spec, err := serviceSpec()
	if err != nil {
		return "", err
	}
	return defaultContainer.StartService(spec)
}

func serviceSpec() (spec container.ServiceOptions, err error) {
	sharedNetworkID, err := defaultContainer.SharedNetworkID()
	if err != nil {
		return container.ServiceOptions{}, err
	}

	_, port, err := xnet.SplitHostPort(viper.GetString(config.APIServerAddress))
	if err != nil {
		return container.ServiceOptions{}, err
	}

	return container.ServiceOptions{
		Namespace: Namespace(),
		Image:     viper.GetString(config.CoreImage),
		Env: container.MapToEnv(map[string]string{
			config.ToEnv(config.MESGPath):  path,
			config.ToEnv(config.LogFormat): viper.GetString(config.LogFormat),
			config.ToEnv(config.LogLevel):  viper.GetString(config.LogLevel),
		}),
		Mounts: []container.Mount{
			{
				Source: dockerSocket,
				Target: dockerSocket,
				Bind:   true,
			},
			{
				Source: volume,
				Target: path,
			},
		},
		Ports: []container.Port{
			{
				Target:    uint32(port),
				Published: uint32(port),
			},
		},
		NetworksID: []string{sharedNetworkID},
	}, nil
}
