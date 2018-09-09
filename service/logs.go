package service

import (
	"io"

	"github.com/mesg-foundation/core/x/xstrings"
)

// Logs holds log streams of dependency.
type Logs struct {
	Dependency string
	Standard   io.ReadCloser
	Error      io.ReadCloser
}

// Logs gives service's logs and applies dependencies filter to filter logs.
// if dependencies has a length of zero all dependency logs will be provided.
func (s *Service) Logs(dependencies ...string) ([]*Logs, error) {
	var (
		logs       []*Logs
		isNoFilter = len(dependencies) == 0
	)
	for _, dep := range s.DependenciesFromService() {
		if isNoFilter || xstrings.SliceContains(dependencies, dep.Name) {
			rstd, rerr, err := dep.Logs()
			if err != nil {
				return nil, err
			}
			logs = append(logs, &Logs{
				Dependency: dep.Name,
				Standard:   rstd,
				Error:      rerr,
			})
		}
	}
	return logs, nil
}
