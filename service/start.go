package service

import (
	"strconv"
	"strings"

	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/x/xnet"
	"github.com/mesg-foundation/core/x/xstructhash"
)

// Start starts the service.
func (s *Service) Start() (serviceIDs []string, err error) {
	status, err := s.Status()
	if err != nil || status == RUNNING {
		return nil, err //TODO: if the service is already running, serviceIDs should be returned.
	}
	// If there is one but not all services running stop to restart all
	if status == PARTIAL {
		if err := s.StopDependencies(); err != nil {
			return nil, err
		}
	}
	networkID, err := s.container.CreateNetwork(s.namespace())
	if err != nil {
		return nil, err
	}
	// BUG: https://github.com/mesg-foundation/core/issues/382
	// After solving this by docker, switch back to deploy in parallel
	serviceIDs = make([]string, len(s.Dependencies))
	for i, dep := range s.Dependencies {
		serviceID, err := dep.Start(networkID)
		if err != nil {
			s.Stop()
			return nil, err
		}
		serviceIDs[i] = serviceID
	}
	return serviceIDs, err
}

// Start starts a dependency container.
func (d *Dependency) Start(networkID string) (containerServiceID string, err error) {
	sharedNetworkID, err := d.service.container.SharedNetworkID()
	if err != nil {
		return "", err
	}
	volumes := d.extractVolumes()
	volumesFrom, err := d.extractVolumesFrom()
	if err != nil {
		return "", err
	}
	c, err := config.Global()
	if err != nil {
		return "", err
	}
	_, port, _ := xnet.SplitHostPort(c.Server.Address)
	endpoint := c.Core.Name + ":" + strconv.Itoa(port)
	return d.service.container.StartService(container.ServiceOptions{
		Namespace: d.namespace(),
		Labels: map[string]string{
			"mesg.service": d.service.Name,
			"mesg.hash":    d.service.Hash,
			"mesg.sid":     d.service.SID,
			"mesg.core":    c.Core.Name,
		},
		Image:   d.Image,
		Args:    d.Args,
		Command: d.Command,
		Env: container.MapToEnv(map[string]string{
			"MESG_TOKEN":        d.service.Hash,
			"MESG_ENDPOINT":     endpoint,
			"MESG_ENDPOINT_TCP": endpoint,
		}),
		Mounts: append(volumes, volumesFrom...),
		Ports:  d.extractPorts(),
		Networks: []container.Network{
			{ID: networkID, Alias: d.Key},
			{ID: sharedNetworkID},
		},
	})
}

func (d *Dependency) extractPorts() []container.Port {
	ports := make([]container.Port, len(d.Ports))
	for i, p := range d.Ports {
		split := strings.Split(p, ":")
		from, _ := strconv.ParseUint(split[0], 10, 64)
		to := from
		if len(split) > 1 {
			to, _ = strconv.ParseUint(split[1], 10, 64)
		}
		ports[i] = container.Port{
			Target:    uint32(to),
			Published: uint32(from),
		}
	}
	return ports
}

// TODO: add test and hack for MkDir in CircleCI
func (d *Dependency) extractVolumes() []container.Mount {
	volumes := make([]container.Mount, 0)
	for _, volume := range d.Volumes {
		volumes = append(volumes, container.Mount{
			Source: volumeKey(d.service, d.Key, volume),
			Target: volume,
		})
	}
	return volumes
}

func (d *Dependency) extractVolumesFrom() ([]container.Mount, error) {
	volumesFrom := make([]container.Mount, 0)
	for _, depName := range d.VolumesFrom {
		dep, err := d.service.getDependency(depName)
		if err != nil {
			return nil, err
		}
		for _, volume := range dep.Volumes {
			volumesFrom = append(volumesFrom, container.Mount{
				Source: volumeKey(d.service, depName, volume),
				Target: volume,
			})
		}
	}
	return volumesFrom, nil
}

// volumeKey creates a key for service's volume based on the sid to make sure that the volume
// will stay the same for different versions of the service.
func volumeKey(s *Service, dependency string, volume string) string {
	return xstructhash.Hash([]string{
		s.SID,
		dependency,
		volume,
	}, 1)
}
