package daemon

import (
	"github.com/mesg-foundation/core/container"
)

// Stop the core docker container
func Stop() (err error) {
	stopped, err := IsStopped()
	if err != nil || stopped == true {
		return
	}
	err = container.StopService(Namespace())
	if err != nil {
		return
	}
	err = container.WaitForContainerStatus(Namespace(), container.STOPPED)
	return
}
