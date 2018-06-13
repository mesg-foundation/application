package service

import (
	"strings"

	"github.com/mesg-foundation/core/utils/hash"
)

// NAMESPACE is the namespace used for the docker services
const eventChannel string = "Event"
const taskChannel string = "Task"
const resultChannel string = "Result"

func (service *Service) namespace() string {
	return strings.Join([]string{
		service.Hash(),
	}, "-")
}

// EventSubscriptionChannel returns the channel to listen for events from this service
func (service *Service) EventSubscriptionChannel() string {
	return hash.Calculate([]string{
		service.namespace(),
		eventChannel,
	})
}

// TaskSubscriptionChannel returns the channel to listen for tasks from this service
func (service *Service) TaskSubscriptionChannel() string {
	return hash.Calculate([]string{
		service.namespace(),
		taskChannel,
	})
}

// ResultSubscriptionChannel returns the channel to listen for tasks from this service
func (service *Service) ResultSubscriptionChannel() string {
	return hash.Calculate([]string{
		service.namespace(),
		resultChannel,
	})
}
