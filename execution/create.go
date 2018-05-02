package execution

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/utils/hash"
)

// Create an execution with a unique ID and put it in the pending list
func Create(service *service.Service, task string, inputs interface{}) (execution *Execution, err error) {
	if !taskExists(service, task) {
		err = errors.New("Task " + task + " doesn't exists in service " + service.Name)
		return
	}
	execution = &Execution{
		Service:   service,
		Inputs:    inputs,
		Task:      task,
		CreatedAt: time.Now(),
	}
	execution.ID, err = generateID(execution)
	if err != nil {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	pendingExecutions[execution.ID] = execution
	log.Println("[PENDING]", task)
	return
}

func taskExists(service *service.Service, name string) (exists bool) {
	exists = false
	for taskName := range service.Tasks {
		if taskName == name {
			exists = true
			break
		}
	}
	return
}

func generateID(execution *Execution) (id string, err error) {
	inputs, err := json.Marshal(execution.Inputs)
	if err != nil {
		return
	}
	id = hash.Calculate([]string{
		execution.CreatedAt.UTC().String(),
		execution.Service.Name,
		execution.Task,
		string(inputs),
	})
	return
}
