package daemon

import (
	"time"

	"github.com/mesg-foundation/core/docker"
)

// IsRunning returns true if the daemon container is running
func IsRunning() (running bool, err error) {
	return docker.IsServiceRunning(Namespace())
}

// IsStopped returns true if the daemon container is stopped
func IsStopped() (stopped bool, err error) {
	return docker.IsServiceStopped(Namespace())
}

// WaitForContainerToRun wait for the Daemon container to run
func WaitForContainerToRun() (wait chan error) {
	return docker.WaitContainerStatus(Namespace(), docker.RUNNING, 5*time.Minute)
}

// WaitForContainerToStop wait for the Daemon container to stop
func WaitForContainerToStop() (wait chan error) {
	return docker.WaitContainerStatus(Namespace(), docker.STOPPED, time.Minute)
}

// WaitForFullStop wait for the daemon container and its shared network to stop
func WaitForFullStop() (wait chan error) {
	wait = make(chan error, 1)
	go func() {
		err := <-WaitForContainerToStop()
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
