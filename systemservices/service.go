package systemservices

import (
	"path/filepath"
	"sync"

	"github.com/docker/docker/pkg/archive"
	"github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/x/xerrors"
	"github.com/mesg-foundation/core/x/xos"
)

// deployServices deploys system services.
func (s *SystemServices) deployServices(services []*systemService) error {
	var (
		// errs are the deployment errors.
		errs xerrors.Errors
		m    sync.Mutex

		wg sync.WaitGroup
	)

	for _, ss := range services {
		wg.Add(1)
		go func(ss *systemService) {
			defer wg.Done()
			s, err := s.deployService(ss.name)
			m.Lock()
			defer m.Unlock()
			if err != nil {
				errs = append(errs, err)
				return
			}
			ss.Service = s
		}(ss)
	}

	wg.Wait()
	return errs.ErrorOrNil()
}

// deployService deploys a system service living in relativePath.
func (s *SystemServices) deployService(relativePath string) (*service.Service, error) {
	path := filepath.Join(s.systemServicesPath, relativePath)
	exists, err := xos.DirExists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &systemServiceNotFoundError{name: relativePath}
	}

	archive, err := archive.TarWithOptions(path, &archive.TarOptions{
		Compression: archive.Gzip,
	})
	if err != nil {
		return nil, err
	}

	service, validationErr, err := s.api.DeployService(archive)
	if err != nil {
		return nil, err
	}
	if validationErr != nil {
		return nil, validationErr
	}
	return service, nil
}

// startService starts the system services.
func (s *SystemServices) startServices(services []*systemService) error {
	var (
		// errs are the service starting errors.
		errs xerrors.Errors
		m    sync.Mutex

		wg sync.WaitGroup
	)

	for _, ss := range services {
		wg.Add(1)
		go func(ss *systemService) {
			defer wg.Done()
			if err := s.api.StartService(ss.ID); err != nil {
				m.Lock()
				defer m.Unlock()
				errs = append(errs, err)
			}
		}(ss)
	}

	wg.Wait()
	return errs.ErrorOrNil()
}

// getServiceID returns the service id of a system service that matches with name.
// name compared with the unique name/relative path of system service.
func (s *SystemServices) getServiceID(services []*systemService, name string) string {
	for _, s := range services {
		if s.name == name {
			return s.ID
		}
	}
	panic("unreachable")
}
