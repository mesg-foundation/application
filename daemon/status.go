package daemon

import (
	"time"

	"github.com/mesg-foundation/core/docker"
)

// IsRunning returns true if the daemon container is running
func IsRunning() (running bool, err error) {
	status, err := docker.ContainerStatus(Namespace())
	if err != nil {
		return
	}
	running = status == docker.RUNNING
	return
}

// WaitForRunning wait for the Daemon container to run
func WaitForRunning() (wait chan error) {
	return docker.WaitContainerStatus(Namespace(), docker.RUNNING, time.Minute)
}

// WaitForStopped wait for the Daemon container to stop
func WaitForStopped() (wait chan error) {
	return docker.WaitContainerStatus(Namespace(), docker.STOPPED, time.Minute)
}

// WaitForFullyStop wait for the daemon container and its shared network to stop
func WaitForFullyStop() (wait chan error) {
	wait = make(chan error, 1)
	go func() {
		err := <-WaitForStopped()
		if err != nil {
			wait <- err
		}
		err = <-docker.WaitNetworkDeletion(NamespaceNetwork(), time.Minute)
		if err != nil {
			wait <- err
		}
		close(wait)
	}()
	return
}
