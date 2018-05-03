package service

import (
	"strings"
	"testing"

	"github.com/stvp/assert"
)

func TestStopRunningService(t *testing.T) {
	service := &Service{
		Name: "TestStopRunningService",
		Dependencies: map[string]*Dependency{
			"test": &Dependency{
				Image: "nginx",
			},
		},
	}
	service.Start()
	err := service.Stop()
	assert.Nil(t, err)
	assert.Equal(t, service.IsStopped(), true)
}

func TestStopNonRunningService(t *testing.T) {
	service := &Service{
		Name: "TestStopNonRunningService",
		Dependencies: map[string]*Dependency{
			"test": &Dependency{
				Image: "nginx",
			},
		},
	}
	err := service.Stop()
	assert.Nil(t, err)
	assert.Equal(t, service.IsStopped(), true)
}

func TestStopDependency(t *testing.T) {
	namespace := strings.Join([]string{NAMESPACE, "TestStopDependency"}, "_")
	name := "test"
	dependency := Dependency{Image: "nginx"}
	network, err := createNetwork(namespace)
	dependency.Start(dependencyDetails{
		namespace:      namespace,
		dependencyName: name,
		serviceName:    "TestStopDependency",
	}, network)
	err = dependency.Stop(namespace, name)
	assert.Nil(t, err)
	assert.Equal(t, dependency.IsStopped(namespace, name), true)
	assert.Equal(t, dependency.IsRunning(namespace, name), false)
	deleteNetwork(namespace)
}

func TestNetworkDeleted(t *testing.T) {
	service := &Service{
		Name: "TestNetworkDeleted",
		Dependencies: map[string]*Dependency{
			"test": &Dependency{
				Image: "nginx",
			},
		},
	}
	service.Start()
	service.Stop()
	network, err := findNetwork(service.namespace())
	assert.Nil(t, err)
	assert.Nil(t, network)
}
