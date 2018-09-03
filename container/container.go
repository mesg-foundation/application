package container

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	docker "github.com/docker/docker/client"
)

// Container provides high level interactions with Docker API for MESG.
type Container struct {
	// client is a Docker client.
	client docker.CommonAPIClient

	// callTimeout is the timeout value for Docker API calls.
	callTimeout time.Duration
}

// Option is a configuration func for Container.
type Option func(*Container)

// New creates a new Container with given options.
func New(options ...Option) (*Container, error) {
	c := &Container{
		callTimeout: 10 * time.Second,
	}
	for _, option := range options {
		option(c)
	}
	var err error
	if c.client == nil {
		c.client, err = docker.NewEnvClient()
		if err != nil {
			return c, err
		}
	}
	c.negotiateAPIVersion()
	if err := c.createSwarmIfNeeded(); err != nil {
		return c, err
	}
	return c, c.createSharedNetworkIfNeeded()
}

// ClientOption receives a client which will be used to interact with Docker API.
func ClientOption(client docker.CommonAPIClient) Option {
	return func(c *Container) {
		c.client = client
	}
}

// TimeoutOption receives d which will be set as a timeout value for Docker API calls.
func TimeoutOption(d time.Duration) Option {
	return func(c *Container) {
		c.callTimeout = d
	}
}

func (c *Container) negotiateAPIVersion() {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	c.client.NegotiateAPIVersion(ctx)
}

func (c *Container) createSwarmIfNeeded() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	info, err := c.client.Info(ctx)
	if err != nil {
		return err
	}
	if info.Swarm.NodeID != "" {
		return nil
	}
	ctx, cancel = context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	_, err = c.client.SwarmInit(ctx, swarm.InitRequest{
		ListenAddr: "0.0.0.0:2377", // https://docs.docker.com/engine/reference/commandline/swarm_init/#usage
	})
	return err
}

// FindContainer returns a docker container.
func (c *Container) FindContainer(namespace []string) (types.ContainerJSON, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	containers, err := c.client.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "label",
			Value: "com.docker.stack.namespace=" + Namespace(namespace),
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

// Status returns the status of a docker container and servcie.
func (c *Container) Status(namespace []string) (StatusType, error) {
	_, err := c.FindContainer(namespace)
	if err != nil && !docker.IsErrNotFound(err) {
		return UNKNOWN, err
	}
	containerFound := !docker.IsErrNotFound(err)
	if containerFound {
		return RUNNING, nil
	}

	_, err = c.FindService(namespace)
	if err != nil && !docker.IsErrNotFound(err) {
		return UNKNOWN, err
	}
	serviceFound := !docker.IsErrNotFound(err)
	if serviceFound {
		return STARTING, nil
	}

	return STOPPED, nil
}
