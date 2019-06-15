package workflowvm

import (
	"fmt"
	"sync"

	"github.com/mesg-foundation/core/workflow"
)

// VM is a virtual machine that runs workflows.
type VM struct {
	// ExecuctionRequests keeps execution requests, caller should read this chan
	// and perform executions.
	ExecuctionRequests chan *Execution

	m               sync.Mutex // protects following.
	workflows       []workflow.Workflow
	executionChains []*executionChain
	closing         bool
}

// New creates a new Workflow VM.
func New() *VM {
	return &VM{
		ExecuctionRequests: make(chan *Execution),
	}
}

// Process processes an event by creating execution requests for matching workflows.
func (v *VM) Process(e Event) {
	v.m.Lock()
	defer v.m.Unlock()
	if v.closing {
		return
	}
	if e.Secret != "" {
		// if secret is non-empty, event is actually a result of previous execution request.
		// generate the next execution request on that chain by using data from this result.
		for i, c := range v.executionChains {
			ok, execReq := c.Next(e)
			// check if we're at the end of the chain, if so, workflow's lifecycle for this
			// chain has completed and it can be removed.
			if c.IsDone() {
				v.executionChains = append(v.executionChains[:i], v.executionChains[i+1:]...)
			}
			if ok {
				v.ExecuctionRequests <- execReq
				break
			}
		}
	}
	// event can also be a start of some new lifecycles for many workflows.
	// create a full trigger spec for the event and find out the matching workflows.
	t := e.createTrigger()
	for _, w := range v.workflows {
		// check if workflow matches with trigger created from event.
		if ok := w.Match(t); ok {
			// create a new chain, register it and create task execution request for
			// its first task.
			c := newChain(w)
			ok, execReq := c.Next(e)
			// chain might have just one task, if so, don't bother to add it chain list.
			if !c.IsDone() {
				v.executionChains = append(v.executionChains, c)
			}
			// check against empty workflow Tasks.
			if ok {
				v.ExecuctionRequests <- execReq
			}
		}
	}
}

// Add adds a workflow to VM.
func (v *VM) Add(w workflow.Workflow) {
	v.m.Lock()
	defer v.m.Unlock()
	if v.closing {
		return
	}
	v.workflows = append(v.workflows, w)
	fmt.Println("vm added:", w.Key)
}

// Remove removes workflow and all of its execution chains from VM.
func (v *VM) Remove(hash string) {
	v.m.Lock()
	defer v.m.Unlock()
	if v.closing {
		return
	}
	for i, w := range v.workflows {
		if w.Hash == hash {
			v.workflows = append(v.workflows[:i], v.workflows[i+1:]...)
			for ii := len(v.executionChains) - 1; i > 0; i-- {
				if v.executionChains[ii].w.Hash == hash {
					v.executionChains = append(v.executionChains[:ii], v.executionChains[ii+1:]...)
				}
			}
			return
		}
	}
}

// Shutdown immediately ends vm and closes ExecutionRequests channel.
func (v *VM) Shutdown() {
	v.m.Lock()
	defer v.m.Unlock()
	v.closing = true
	close(v.ExecuctionRequests)
}

// Event represents a new event/result input to the workflow VM.
type Event struct {
	// InstanceHash that event is emitted from.
	InstanceHash string

	// ParentHash of event if there is.
	ParentHash []byte

	// Key of event.
	Key string

	// TaskKey of event when event is emitted as a response to an execution.
	TaskKey string

	// Data is the payload of event.
	Data map[string]interface{}

	// Secret is generated by the Workflow VM for each execution request in a chain to check which
	// Event belongs to which Execution request's result.
	Secret string
}

// createTrigger creates a new trigger schema from event which is used to match the
// event with registered workflows in VM.
func (e Event) createTrigger() workflow.Trigger {
	return workflow.Trigger{
		EventKey:     e.Key,
		InstanceHash: e.InstanceHash,
		Filter: workflow.Filter{
			TaskKey: e.TaskKey,
		},
	}
}

// Execution is an execution request that should be performed by the caller.
// Each received Event from the Process() may cause multiple Execution requests
// to be emitted by the Workflow VM.
type Execution struct {
	// InstanceHash of instance that execution should be made on.
	InstanceHash string

	// ParentHash is the parent execution hash if execution has one.
	ParentHash []byte

	// TaskKey of execution.
	TaskKey string

	// Inputs of the execution.
	Inputs map[string]interface{}

	// Secret is generated by the Workflow VM for each execution request in a chain to check which
	// Event belongs to which Execution request's result.
	Secret string
}
