package container

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	docker "github.com/docker/docker/client"
	"github.com/mesg-foundation/core/config"
)

var (
	errSwarmNotInit = errors.New(`docker swarm is not initialized. run "docker swarm init" and try again`)
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
	DeleteVolume(name string) error
}

// DockerContainer provides high level interactions with Docker API for MESG.
type DockerContainer struct {
	// client is a Docker client.
	client docker.CommonAPIClient

	// callTimeout is the timeout value for Docker API calls.
	callTimeout time.Duration

	config *config.Config
}

// Option is a configuration func for Container.
type Option func(*DockerContainer)

// New creates a new Container with given options.
func New(options ...Option) (*DockerContainer, error) {
	c := &DockerContainer{
		callTimeout: 10 * time.Second,
	}
	for _, option := range options {
		option(c)
	}
	var err error
	cfg, err := config.Global()
	if err != nil {
		return nil, err
	}
	c.config = cfg
	if c.client == nil {
		c.client, err = docker.NewClientWithOpts(docker.FromEnv)
		if err != nil {
			return c, err
		}
	}
	c.negotiateAPIVersion()
	if err := c.isSwarmInit(); err != nil {
		return c, err
	}
	return c, c.createSharedNetworkIfNeeded()
}

// ClientOption receives a client which will be used to interact with Docker API.
func ClientOption(client docker.CommonAPIClient) Option {
	return func(c *DockerContainer) {
		c.client = client
	}
}

// TimeoutOption receives d which will be set as a timeout value for Docker API calls.
func TimeoutOption(d time.Duration) Option {
	return func(c *DockerContainer) {
		c.callTimeout = d
	}
}

func (c *DockerContainer) negotiateAPIVersion() {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	c.client.NegotiateAPIVersion(ctx)
}

func (c *DockerContainer) isSwarmInit() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	info, err := c.client.Info(ctx)
	if err != nil {
		return err
	}
	if info.Swarm.NodeID == "" {
		return errSwarmNotInit
	}
	return nil
}

// FindContainer returns a docker container.
func (c *DockerContainer) FindContainer(namespace []string) (types.ContainerJSON, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	containers, err := c.client.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "label",
			Value: "com.docker.stack.namespace=" + c.Namespace(namespace),
		}),
		Limit: 1,
	})
	if err != nil {
		return types.ContainerJSON{}, err
	}
	containerID := ""
	if len(containers) == 1 {
		containerID = containers[0].ID
	}
	ctx, cancel = context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	return c.client.ContainerInspect(ctx, containerID)
}

// Status returns the status of the container based on the docker container and docker service.
// if any error occurs during the status check, status will be shown as UNKNOWN.
// otherwise the following rules will be applied to determine a status:
//  - RUNNING: when the container is running in docker regardless of the status of the service.
//  - STARTING: when the service is running but the container is not yet started.
//  - STOPPED: when the container and the service is not running in docker.
func (c *DockerContainer) Status(namespace []string) (StatusType, error) {
	container, err := c.containerExists(namespace)
	if err != nil {
		return UNKNOWN, err
	}
	service, err := c.serviceExists(namespace)
	if err != nil {
		return UNKNOWN, err
	}

	statuses := []struct {
		container bool
		service   bool
		status    StatusType
	}{
		{service: true, container: true, status: RUNNING},
		{service: true, container: false, status: STARTING},
		{service: false, container: true, status: RUNNING}, // This is actually stopping
		{service: false, container: false, status: STOPPED},
	}

	for _, s := range statuses {
		if s.container == container && s.service == service {
			return s.status, nil
		}
	}
	return UNKNOWN, nil // This should never be reached but it's better than a panic :)
}

// containerExists checks if container with namespace can be found.
func (c *DockerContainer) containerExists(namespace []string) (bool, error) {
	_, err := c.FindContainer(namespace)
	return presenceHandling(err)
}

// serviceExists checks if corresponding container for service namespace can be found.
func (c *DockerContainer) serviceExists(namespace []string) (bool, error) {
	_, err := c.FindService(namespace)
	return presenceHandling(err)
}

// presenceHandling checks err to see if it's a Docker NotFound error and if not
// it'll return the err back.
func presenceHandling(err error) (bool, error) {
	if err != nil && !docker.IsErrNotFound(err) {
		return false, err
	}
	return !docker.IsErrNotFound(err), nil
}
