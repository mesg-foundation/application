package api

import (
	"fmt"

	"github.com/mesg-foundation/core/execution"
	"github.com/mesg-foundation/core/pubsub"
	"github.com/mesg-foundation/core/service"
	uuid "github.com/satori/go.uuid"
)

// ExecuteTask executes a task tasKey with inputData and tags for service serviceID.
func (a *API) ExecuteTask(serviceID, taskKey string, inputData map[string]interface{},
	tags []string) (executionID string, err error) {
	return newTaskExecutor(a).Execute(serviceID, taskKey, inputData, tags)
}

// taskExecutor provides functionalities to execute a MESG task.
type taskExecutor struct {
	api *API
}

// newTaskExecutor creates a new taskExecutor with given api.
func newTaskExecutor(api *API) *taskExecutor {
	return &taskExecutor{
		api: api,
	}
}

// ExecuteTask executes a task tasKey with inputData and tags for service serviceID.
func (e *taskExecutor) Execute(serviceID, taskKey string, inputData map[string]interface{},
	tags []string) (executionID string, err error) {
	s, err := e.api.db.Get(serviceID)
	if err != nil {
		return "", err
	}
	s, err = service.FromService(s, service.ContainerOption(e.api.container))
	if err != nil {
		return "", err
	}
	if err := e.checkServiceStatus(s); err != nil {
		return "", err
	}
	return e.execute(s, uuid.NewV4().String(), taskKey, inputData, tags)
}

// checkServiceStatus checks service status. A task should be executed only if
// task's service is running.
func (e *taskExecutor) checkServiceStatus(s *service.Service) error {
	status, err := s.Status()
	if err != nil {
		return err
	}
	if status != service.RUNNING {
		return &NotRunningServiceError{ServiceID: s.Hash}
	}
	return nil
}

// execute executes task.
func (e *taskExecutor) execute(s *service.Service, eventID string, taskKey string, taskInputs map[string]interface{}, tags []string) (executionID string, err error) {
	exec, err := execution.New(s, eventID, taskKey, taskInputs, tags)
	if err != nil {
		return "", err
	}
	if err := exec.Execute(); err != nil {
		return "", err
	}
	if err = e.api.execDB.Save(exec); err != nil {
		return "", err
	}
	go pubsub.Publish(s.TaskSubscriptionChannel(), exec)
	return exec.ID, nil
}

// NotRunningServiceError is an error returned when the service is not running that
// a task needed to be executed on.
type NotRunningServiceError struct {
	ServiceID string
}

func (e *NotRunningServiceError) Error() string {
	return fmt.Sprintf("Service %q is not running", e.ServiceID)
}
