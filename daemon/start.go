package daemon

import (
	"path/filepath"

	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/x/xnet"
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
	c, err := config.Global()
	if err != nil {
		return container.ServiceOptions{}, err
	}
	_, port, err := xnet.SplitHostPort(c.Server.Address)
	if err != nil {
		return container.ServiceOptions{}, err
	}
	return container.ServiceOptions{
		Namespace: []string{},
		Image:     c.Core.Image,
		Env:       container.MapToEnv(c.DaemonEnv()),
		Mounts: []container.Mount{
			{
				Source: c.Docker.Socket,
				Target: c.Docker.Socket,
				Bind:   true,
			},
			{
				Source: c.Core.Path,
				Target: c.Docker.Core.Path,
				Bind:   true,
			},
			{
				Source: filepath.Join(c.Core.Path, c.SystemServices.RelativePath),
				Target: filepath.Join(c.Docker.Core.Path, c.SystemServices.RelativePath),
				Bind:   true,
			},
		},
		Ports: []container.Port{
			{
				Target:    uint32(port),
				Published: uint32(port),
			},
		},
		Networks: []container.Network{
			{ID: sharedNetworkID},
		},
	}, nil
}
