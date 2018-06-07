package daemon

import (
	"io"

	"github.com/mesg-foundation/core/container"
)

// Logs return the core's docker service logs
func Logs() (reader io.ReadCloser, err error) {
	return container.ServiceLogs(Namespace())
}
