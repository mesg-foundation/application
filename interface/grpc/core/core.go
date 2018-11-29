package core

import (
	"github.com/mesg-foundation/core/protobuf/coreapi"
	service "github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/utils/workflowparser"
)

func toProtoServices(ss []*service.Service) []*coreapi.Service {
	services := make([]*coreapi.Service, len(ss))
	for i, s := range ss {
		services[i] = toProtoService(s)
	}
	return services
}

func toProtoService(s *service.Service) *coreapi.Service {
	return &coreapi.Service{
		ID:           s.ID,
		Name:         s.Name,
		Description:  s.Description,
		Repository:   s.Repository,
		Tasks:        toProtoTasks(s.Tasks),
		Events:       toProtoEvents(s.Events),
		Dependencies: toProtoDependencies(s.Dependencies),
	}
}

func toProtoServiceStatusType(s service.StatusType) coreapi.Service_Status {
	switch s {
	default:
		return coreapi.Service_UNKNOWN
	case service.STOPPED:
		return coreapi.Service_STOPPED
	case service.STARTING:
		return coreapi.Service_STARTING
	case service.PARTIAL:
		return coreapi.Service_PARTIAL
	case service.RUNNING:
		return coreapi.Service_RUNNING
	}
}

func toProtoTasks(tasks []*service.Task) []*coreapi.Task {
	ts := make([]*coreapi.Task, len(tasks))
	for i, task := range tasks {
		t := &coreapi.Task{
			Key:         task.Key,
			Name:        task.Name,
			Description: task.Description,
			Inputs:      toProtoParameters(task.Inputs),
			Outputs:     []*coreapi.Output{},
		}
		for _, output := range task.Outputs {
			o := &coreapi.Output{
				Key:         output.Key,
				Name:        output.Name,
				Description: output.Description,
				Data:        toProtoParameters(output.Data),
			}
			t.Outputs = append(t.Outputs, o)
		}
		ts[i] = t
	}
	return ts
}

func toProtoEvents(events []*service.Event) []*coreapi.Event {
	es := make([]*coreapi.Event, len(events))
	for i, event := range events {
		es[i] = &coreapi.Event{
			Key:         event.Key,
			Name:        event.Name,
			Description: event.Description,
			Data:        toProtoParameters(event.Data),
		}
	}
	return es
}

func toProtoParameters(params []*service.Parameter) []*coreapi.Parameter {
	ps := make([]*coreapi.Parameter, len(params))
	for i, param := range params {
		ps[i] = &coreapi.Parameter{
			Key:         param.Key,
			Name:        param.Name,
			Description: param.Description,
			Type:        param.Type,
			Optional:    param.Optional,
		}
	}
	return ps
}

func toProtoDependency(dep *service.Dependency) *coreapi.Dependency {
	if dep == nil {
		return nil
	}
	return &coreapi.Dependency{
		Key:         dep.Key,
		Image:       dep.Image,
		Volumes:     dep.Volumes,
		Volumesfrom: dep.VolumesFrom,
		Ports:       dep.Ports,
		Command:     dep.Command,
	}
}

func toProtoDependencies(deps []*service.Dependency) []*coreapi.Dependency {
	ds := make([]*coreapi.Dependency, len(deps))
	for i, dep := range deps {
		ds[i] = toProtoDependency(dep)
	}
	return ds
}

func toWorkflowDefinition(_ *coreapi.CreateWorkflowRequest_WorkflowDefinition) workflowparser.WorkflowDefinition {
	return workflowparser.WorkflowDefinition{}
}
