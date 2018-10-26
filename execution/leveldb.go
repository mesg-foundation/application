package execution

import (
	"crypto/sha1"
	"encoding/json"
	"time"

	"github.com/cnf/structhash"
	"github.com/mesg-foundation/core/service"
	"github.com/syndtr/goleveldb/leveldb"
)

// LevelDB is a concrete implementation of the DB interface
type LevelDB struct {
	db *leveldb.DB
}

// New creates a new DB instance
func New(path string) (DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &LevelDB{db: db}, nil
}

// Create a record in the database to store this execution and returns the id
// returns an error if any problem happen with the database
// returns an error if inputs are invalid
func (db *LevelDB) Create(task service.Task, inputs map[string]interface{}, tags []string) (*Execution, error) {
	if err := task.RequireInputs(inputs); err != nil {
		return nil, err
	}
	return db.save(&Execution{
		Inputs:    inputs,
		Task:      task,
		Tags:      tags,
		CreatedAt: time.Now(),
		Status:    Created,
	})
}

// Find the execution based on an executionID, returns an error if not found
func (db *LevelDB) Find(executionID []byte) (*Execution, error) {
	data, err := db.db.Get(executionID, nil)
	if err != nil {
		return nil, err
	}
	var execution Execution
	err = json.Unmarshal(data, &execution)
	return &execution, err
}

// Execute a given execution
// Returns an error if the execution doesn't exists in the database
// Returns an error if the status of the execution is different of `Created`
func (db *LevelDB) Execute(executionID []byte) (*Execution, error) {
	e, err := db.Find(executionID)
	if err != nil {
		return nil, err
	}
	if e.Status != Created {
		return nil, StatusError{}
	}
	e.ExecutedAt = time.Now()
	e.Status = InProgress
	return db.save(e)
}

// Complete verifies the output associated to the execution and save this to the database
// Returns an error if the executionID doesn't exists
// Returns an error if the execution is not `InProgress`
// Returns an error if the `outputKey` or `outputData` are not valid
func (db *LevelDB) Complete(executionID []byte, outputKey string, outputData map[string]interface{}) (*Execution, error) {
	e, err := db.Find(executionID)
	if err != nil {
		return nil, err
	}
	if e.Status != InProgress {
		return nil, StatusError{}
	}
	output, err := e.Task.GetOutput(outputKey)
	if err != nil {
		return nil, err
	}
	if err := output.RequireData(outputData); err != nil {
		return nil, err
	}

	e.ExecutionDuration = time.Since(e.ExecutedAt)
	e.Output = outputKey
	e.OutputData = outputData
	e.Status = Completed

	return db.save(e)
}

func (db *LevelDB) save(execution *Execution) (*Execution, error) {
	h := sha1.New()
	h.Write(structhash.Dump(execution, 1))
	id := h.Sum(nil)
	execution.ID = id
	data, err := json.Marshal(execution)
	if err != nil {
		return nil, err
	}
	err = db.db.Put(id, data, nil)
	return execution, nil
}
