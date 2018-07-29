package container

import (
	"context"

	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
)

// CreateNetwork creates a Docker Network with a namespace.
func (c *Container) CreateNetwork(namespace []string) (id string, err error) {
	network, err := c.FindNetwork(namespace)
	if err != nil && !docker.IsErrNotFound(err) {
		return "", err
	}
	if network.ID != "" {
		return network.ID, nil
	}
	namespaceFlat := Namespace(namespace)
	response, err := c.client.NetworkCreate(context.Background(), namespaceFlat, types.NetworkCreate{
		CheckDuplicate: true, // Cannot have 2 network with the same name
		Driver:         "overlay",
		Labels: map[string]string{
			"com.docker.stack.namespace": namespaceFlat,
		},
	})
	if err != nil {
		return "", err
	}
	return response.ID, nil
}

// DeleteNetwork deletes a Docker Network associated with a namespace.
func (c *Container) DeleteNetwork(namespace []string) error {
	network, err := c.FindNetwork(namespace)
	if docker.IsErrNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return c.client.NetworkRemove(context.Background(), network.ID)
}

// FindNetwork finds a Docker Network by a namespace. If no network is found, an error is returned.
func (c *Container) FindNetwork(namespace []string) (types.NetworkResource, error) {
	return c.client.NetworkInspect(context.Background(), Namespace(namespace), types.NetworkInspectOptions{})
}
