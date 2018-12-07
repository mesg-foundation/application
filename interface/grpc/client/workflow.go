package client

import (
	"context"
	"errors"
	"strings"

	"github.com/mesg-foundation/core/protobuf/coreapi"
)

// Start is the function to start the workflow
func (wf *Workflow) Start() (err error) {
	switch {
	case wf.Execute == nil:
		err = errors.New("a workflow needs a task")
	case wf.OnEvent == nil && wf.OnResult == nil:
		err = errors.New("a workflow needs an event OnEvent or OnResult")
	}
	if err != nil {
		return
	}

	wf.client, err = getClient()
	if err != nil {
		return
	}
	err = startServices(wf)
	if err != nil {
		return
	}
	if wf.OnEvent != nil {
		err = listenEvents(wf)
	} else {
		err = listenResults(wf)
	}
	return
}

// Stop will stop all the services in your workflow
func (wf *Workflow) Stop() (err error) {
	err = stopServices(wf)
	return
}

func listenEvents(wf *Workflow) (err error) {
	if wf.OnEvent.Name == "" {
		err = errors.New("event's Name should be defined (you can use * to react to any event)")
		return
	}
	stream, err := wf.client.ListenEvent(context.Background(), &coreapi.ListenEventRequest{
		ServiceID: wf.OnEvent.ServiceID,
	})
	if err != nil {
		return
	}

	for {
		var data *coreapi.EventData
		data, err = stream.Recv()
		if err != nil {
			break
		}
		if wf.validEvent(data) {
			err = wf.Execute.processEvent(wf, data)
			if err != nil {
				break
			}
		}
	}
	return
}

func (wf *Workflow) validEvent(data *coreapi.EventData) bool {
	if strings.Compare(wf.OnEvent.Name, "*") == 0 {
		return true
	}
	return strings.Compare(wf.OnEvent.Name, data.EventKey) == 0
}

func listenResults(wf *Workflow) (err error) {
	if wf.OnResult.Name == "" || wf.OnResult.Output == "" {
		err = errors.New("result's Name and Output should be defined (you can use * to react to any result)")
		return
	}
	stream, err := wf.client.ListenResult(context.Background(), &coreapi.ListenResultRequest{
		ServiceID: wf.OnResult.ServiceID,
	})
	if err != nil {
		return
	}

	for {
		var data *coreapi.ResultData
		data, err = stream.Recv()
		if err != nil {
			break
		}
		if wf.validResult(data) {
			err = wf.Execute.processResult(wf, data)
			if err != nil {
				break
			}
		}
	}
	return
}

func (wf *Workflow) validResult(data *coreapi.ResultData) bool {
	validName := strings.Compare(wf.OnResult.Name, "*") == 0
	if !validName {
		validName = strings.Compare(wf.OnResult.Name, data.TaskKey) == 0
	}
	validOutput := strings.Compare(wf.OnResult.Output, "*") == 0
	if !validOutput {
		validOutput = strings.Compare(wf.OnResult.Output, data.OutputKey) == 0
	}
	return validName && validOutput
}
