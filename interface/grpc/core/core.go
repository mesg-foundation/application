package core

import service "github.com/mesg-foundation/core/service"

func toProtoServices(ss []*service.Service) []*Service {
	services := make([]*Service, 0)
	for _, s := range ss {
		services = append(services, toProtoService(s))
	}
	return services
}

func toProtoService(s *service.Service) *Service {
	sv := &Service{
		ID:            s.ID,
		Name:          s.Name,
		Description:   s.Description,
		Tasks:         map[string]*Task{},
		Dependencies:  toProtoDependencies(s.Dependencies),
		Configuration: toProtoDependency(s.Configuration),
		Repository:    s.Repository,
	}

	for taskKey, task := range s.Tasks {
		t := &Task{
			Key:         task.Key,
			Name:        task.Name,
			Description: task.Description,
			ServiceName: task.ServiceName,
			Inputs:      toProtoParameters(task.Inputs),
			Outputs:     map[string]*Output{},
		}
		for outputKey, output := range task.Outputs {
			t.Outputs[outputKey] = &Output{
				Key:         output.Key,
				Name:        output.Name,
				Description: output.Description,
				TaskKey:     output.TaskKey,
				ServiceName: output.ServiceName,
				Data:        toProtoParameters(output.Data),
			}
		}
		sv.Tasks[taskKey] = t
	}

	return sv
}

func toProtoParameters(params map[string]*service.Parameter) map[string]*Parameter {
	ps := make(map[string]*Parameter, 0)
	for key, param := range params {
		ps[key] = &Parameter{
			Name:        param.Name,
			Description: param.Description,
			Type:        param.Type,
			Optional:    param.Optional,
		}
	}
	return ps
}

func toProtoDependency(dep *service.Dependency) *Dependency {
	return &Dependency{
		Image:       dep.Image,
		Volumes:     dep.Volumes,
		Volumesfrom: dep.VolumesFrom,
		Ports:       dep.Ports,
		Command:     dep.Command,
	}
}

func toProtoDependencies(deps map[string]*service.Dependency) map[string]*Dependency {
	ds := make(map[string]*Dependency, 0)
	for key, dep := range deps {
		ds[key] = toProtoDependency(dep)
	}
	return ds
}
