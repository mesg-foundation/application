package service

import (
	"errors"
	"strings"
	"sync"

	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/container"
	"github.com/spf13/viper"
)

// Start a service
func (service *Service) Start() (serviceIDs []string, err error) {
	if service.IsRunning() {
		return
	}
	// If there is one but not all services running stop to restart all
	if service.IsPartiallyRunning() {
		err = service.StopDependencies()
		if err != nil {
			return
		}
	}
	networkID, err := container.CreateNetwork([]string{service.namespace()})
	if err != nil {
		return
	}
	serviceIDs = make([]string, len(service.GetDependencies()))
	var mutex sync.Mutex
	i := 0
	var wg sync.WaitGroup
	for name, dependency := range service.GetDependencies() {
		d := dependencyDetails{
			namespace:      service.namespace(),
			dependencyName: name,
			serviceName:    service.Name,
			serviceHash:    service.Hash(),
		}
		wg.Add(1)
		go func(service *Service, dep *Dependency, d dependencyDetails, name string, i int) {
			defer wg.Done()
			serviceID, errStart := dep.Start(service, d, networkID)
			mutex.Lock()
			defer mutex.Unlock()
			serviceIDs[i] = serviceID
			if errStart != nil && err == nil {
				err = errStart
			}
		}(service, dependency, d, name, i)
		i++
	}
	wg.Wait()
	// Disgrasfully close the service because there is an error
	if err != nil {
		service.Stop()
	}
	return
}

type dependencyDetails struct {
	namespace      string
	dependencyName string
	serviceName    string
	serviceHash    string
}

// Start will start a dependency container
func (dependency *Dependency) Start(service *Service, details dependencyDetails, networkID string) (serviceID string, err error) {
	if networkID == "" {
		panic(errors.New("Network ID should never be null"))
	}
	sharedNetworkID, err := container.SharedNetworkID()
	if err != nil {
		return
	}
	namespace := []string{details.namespace, details.dependencyName} //TODO: refacto namespace
	serviceID, err = container.StartService(container.ServiceOptions{
		Namespace: namespace,
		Labels: map[string]string{
			"mesg.service": details.serviceName,
			"mesg.hash":    details.serviceHash,
		},
		Image: dependency.Image,
		Args:  strings.Fields(dependency.Command),
		Env: []string{
			"MESG_TOKEN=" + details.serviceHash,
			"MESG_ENDPOINT=" + viper.GetString(config.APIServiceTargetSocket),
			"MESG_ENDPOINT_TCP=mesg-core:50052", // TODO: should get this from daemon namespace and config
		},
		Mounts: append(extractVolumes(service, dependency, details), container.Mount{
			Source: viper.GetString(config.APIServiceSocketPath),
			Target: viper.GetString(config.APIServiceTargetPath),
		}),
		Ports:      extractPorts(dependency),
		NetworksID: []string{networkID, sharedNetworkID},
	})
	if err != nil {
		return
	}
	err = container.WaitForStatus(namespace, container.RUNNING)
	return
}
