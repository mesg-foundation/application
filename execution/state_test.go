package execution

import (
	"testing"

	"github.com/mesg-foundation/core/service"
	"github.com/stvp/assert"
)

func TestMoveFromPendingToInProgress(t *testing.T) {
	s := service.Service{
		Name: "TestMoveFromPendingToInProgress",
		Tasks: map[string]*service.Task{
			"test": &service.Task{},
		},
	}
	var inputs map[string]interface{}
	exec, _ := Create(&s, "test", inputs)
	err := exec.moveFromPendingToInProgress()
	assert.Equal(t, inProgressExecutions[exec.ID], exec)
	assert.Nil(t, pendingExecutions[exec.ID])
	assert.Nil(t, err)
}

func TestMoveFromPendingToInProgressNonExistingTask(t *testing.T) {
	exec := Execution{ID: "test"}
	err := exec.moveFromPendingToInProgress()
	assert.NotNil(t, err)
	assert.Nil(t, pendingExecutions[exec.ID])
}

func TestMoveFromInProgressToCompleted(t *testing.T) {
	s := service.Service{
		Name: "TestMoveFromInProgressToCompleted",
		Tasks: map[string]*service.Task{
			"test": &service.Task{},
		},
	}
	var inputs map[string]interface{}
	exec, _ := Create(&s, "test", inputs)
	exec.moveFromPendingToInProgress()
	err := exec.moveFromInProgressToProcessed()
	assert.Equal(t, processedExecutions[exec.ID], exec)
	assert.Nil(t, inProgressExecutions[exec.ID])
	assert.Nil(t, err)
}

func TestMoveFromInProgressToCompletedNonExistingTask(t *testing.T) {
	s := service.Service{
		Name: "TestMoveFromInProgressToCompletedNonExistingTask",
		Tasks: map[string]*service.Task{
			"test": &service.Task{},
		},
	}
	var inputs map[string]interface{}
	exec, _ := Create(&s, "test", inputs)
	err := exec.moveFromInProgressToProcessed()
	assert.NotNil(t, err)
	assert.Nil(t, inProgressExecutions[exec.ID])
}
