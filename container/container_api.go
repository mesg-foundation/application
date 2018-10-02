package container

import (
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
)

// Container describes the API of container package.
type Container interface {
	Build(path string) (tag string, err error)
	CreateNetwork(namespace []string) (id string, err error)
	DeleteNetwork(namespace []string, event EventType) error
	FindContainer(namespace []string) (types.ContainerJSON, error)
	FindNetwork(namespace []string) (types.NetworkResource, error)
	FindService(namespace []string) (swarm.Service, error)
	ListServices(labels ...string) ([]swarm.Service, error)
	ListTasks(namespace []string) ([]swarm.Task, error)
	Namespace(ss []string) string
	ServiceLogs(namespace []string) (io.ReadCloser, error)
	SharedNetworkID() (networkID string, err error)
	StartService(options ServiceOptions) (serviceID string, err error)
	Status(namespace []string) (StatusType, error)
	StopService(namespace []string) (err error)
	TasksError(namespace []string) ([]string, error)
}
