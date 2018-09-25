package service

import (
	"github.com/mesg-foundation/core/container"
)

// StatusType of the service.
type StatusType uint

// Possible statuses for service.
const (
	UNKNOWN StatusType = iota
	STOPPED
	STARTING
	PARTIAL
	RUNNING
)

func (s StatusType) String() string {
	switch s {
	case STOPPED:
		return "STOPPED"
	case STARTING:
		return "STARTING"
	case PARTIAL:
		return "PARTIAL"
	case RUNNING:
		return "RUNNING"
	default:
		return "UNKNOWN"
	}
}

var containerStatusTypeMappings = map[container.StatusType]StatusType{
	container.UNKNOWN:  UNKNOWN,
	container.STOPPED:  STOPPED,
	container.STARTING: STARTING,
	container.RUNNING:  RUNNING,
}

// Status returns StatusType of all dependency.
func (s *Service) Status() (StatusType, error) {
	statuses := make(map[container.StatusType]bool)
	for _, dep := range s.Dependencies {
		status, err := dep.Status()
		if err != nil {
			return UNKNOWN, err
		}
		statuses[status] = true
	}

	switch len(statuses) {
	case 0:
		return STOPPED, nil
	case 1:
		for status := range statuses {
			return containerStatusTypeMappings[status], nil
		}
	default:
		return PARTIAL, nil
	}
	panic("not reached")
}

// Status returns StatusType of dependency's container.
func (d *Dependency) Status() (container.StatusType, error) {
	return d.service.docker.Status(d.namespace())
}
