package execution

import (
	"encoding/hex"
	"time"

	"github.com/cnf/structhash"
	"github.com/mesg-foundation/core/service"
)

// Status stores the state of an execution
type Status int

// Status for an execution
// Created    => The execution is created but not yet processed
// InProgress => The execution is being processed
// Completed  => The execution is completed
const (
	Created Status = iota
	InProgress
	Completed
	Failed
)

func (s Status) String() (r string) {
	switch s {
	case Created:
		r = "created"
	case InProgress:
		r = "in progress"
	case Completed:
		r = "completed"
	case Failed:
		r = "failed"
	}
	return r
}

// hashVersion is version for calculate execution hash.
const hashVersion = 1

// Execution stores all informations about executions.
type Execution struct {
	ID                string                 `hash:"-"`
	EventID           string                 `hash:"name:eventID"`
	Status            Status                 `hash:"-"`
	Service           *service.Service       `hash:"name:service"`
	TaskKey           string                 `hash:"name:taskKey"`
	Tags              []string               `hash:"name:tags"`
	Inputs            map[string]interface{} `hash:"name:inputs"`
	OutputKey         string                 `hash:"-"`
	OutputData        map[string]interface{} `hash:"-"`
	Error             string                 `hash:"-"`
	CreatedAt         time.Time              `hash:"-"`
	ExecutedAt        time.Time              `hash:"-"`
	ExecutionDuration time.Duration          `hash:"-"`
}

// New returns a new execution. It returns an error if inputs are invalid.
func New(service *service.Service, eventID string, taskKey string, inputs map[string]interface{}, tags []string) *Execution {
	exec := &Execution{
		EventID:   eventID,
		Service:   service,
		Inputs:    inputs,
		TaskKey:   taskKey,
		Tags:      tags,
		CreatedAt: time.Now(),
		Status:    Created,
	}

	h := structhash.Sha1(exec, hashVersion)
	exec.ID = hex.EncodeToString(h[:])
	return exec
}

// Execute changes executions status to in progres and update its execute time.
// It returns an error if the status is different then Created.
func (execution *Execution) Execute() error {
	if execution.Status != Created {
		return StatusError{
			ExpectedStatus: Created,
			ActualStatus:   execution.Status,
		}
	}
	execution.ExecutedAt = time.Now()
	execution.Status = InProgress
	return nil
}

// Complete changes execution status to completed. It verifies the output.
// It returns an error if the status is different then InProgress or verification fails.
func (execution *Execution) Complete(outputKey string, outputData map[string]interface{}) error {
	if execution.Status != InProgress {
		return StatusError{
			ExpectedStatus: InProgress,
			ActualStatus:   execution.Status,
		}
	}

	execution.ExecutionDuration = time.Since(execution.ExecutedAt)
	execution.OutputKey = outputKey
	execution.OutputData = outputData
	execution.Status = Completed
	return nil
}

// Failed changes execution status to failed and puts error information to execution.
// It returns an error if the status is different then InProgress.
func (execution *Execution) Failed(err error) error {
	if execution.Status != InProgress {
		return StatusError{
			ExpectedStatus: InProgress,
			ActualStatus:   execution.Status,
		}
	}

	execution.Error = err.Error()
	execution.ExecutionDuration = time.Since(execution.ExecutedAt)
	execution.Status = Failed
	return nil
}
