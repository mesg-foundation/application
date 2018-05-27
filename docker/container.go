package docker

import (
	"context"
	"errors"
	"fmt"
	"time"

	godocker "github.com/fsouza/go-dockerclient"
)

// FindContainerStrict returns a docker container if exist. Return error if not found.
func FindContainerStrict(namespace []string) (container *godocker.APIContainers, err error) {
	container, err = FindContainer(namespace)
	if err != nil {
		return
	}
	if container == nil {
		err = errors.New("Container " + Namespace(namespace) + " not found")
	}
	return
}

// FindContainer returns a docker container if exist
func FindContainer(namespace []string) (container *godocker.APIContainers, err error) {
	client, err := Client()
	if err != nil {
		return
	}
	containers, err := client.ListContainers(godocker.ListContainersOptions{
		Context: context.Background(),
		Limit:   1,
		Filters: map[string][]string{
			"label": []string{
				"com.docker.stack.namespace=" + Namespace(namespace),
			},
		},
	})
	if err != nil {
		return
	}
	fmt.Println("containers", containers)
	if len(containers) == 1 {
		container = &containers[0]
	}
	return
}

// ContainerStatus returns the status of a docker container
func ContainerStatus(namespace []string) (status StatusType, err error) {
	container, err := FindContainer(namespace)
	if err != nil {
		return
	}
	status = STOPPED
	if container != nil {
		fmt.Println("container.State", container.State)
	} else {
		fmt.Println("container is nil")
	}
	if container != nil && container.State == "running" {
		status = RUNNING
	}
	return
}

// WaitContainerStatus wait for the container to have the provided status until it reach the timeout
func WaitContainerStatus(namespace []string, status StatusType, timeout time.Duration) (wait chan error) {
	start := time.Now()
	wait = make(chan error, 1)
	go func() {
		for {
			currentStatus, err := ContainerStatus(namespace)
			if err != nil {
				wait <- err
				return
			}
			if currentStatus == status {
				close(wait)
				return
			}
			diff := time.Now().Sub(start)
			if diff.Nanoseconds() >= int64(timeout) {
				wait <- errors.New("Wait too long for the container, timeout reached")
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return
}
