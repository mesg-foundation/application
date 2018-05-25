package client

// Workflow is a struct that contains all the details of
// a workflow. A workflow contains an event source and
// triggers one or multiple tasks. The workflow is what
// is created on the **when**
type Workflow struct {
	Event *Event
	Tasks []*Task
}

// Task is a struct that contains the details of a task
// a task should be associated to a workflow.
// A task is corresponding to the **then** in a workflow
type Task struct {
	Service string
	Name    string
	Inputs  func(interface{}) interface{}
}

// Event is a struct that contains all the informations
// to start a workflow. This is the **when** in the
// workflow
type Event struct {
	Service string
	Name    string
}
