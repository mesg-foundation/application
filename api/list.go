package api

import (
	"fmt"

	"github.com/mesg-foundation/core/service"
)

// ListServicesFilter is a filter func for listing services.
type ListServicesFilter func(*serviceLister)

// ListRunningServicesFilter returns an option to filter services that are
// in the starting process or already started.
func ListRunningServicesFilter() ListServicesFilter {
	return func(l *serviceLister) {
		l.filterActive = true
	}
}

// ListServices returns services matches with filters.
func (a *API) ListServices(filters ...ListServicesFilter) ([]*service.Service, error) {
	return newServiceLister(a, filters...).List()
}

// serviceLister provides functionalities to list MESG services.
type serviceLister struct {
	api          *API
	filterActive bool
}

// newServiceLister creates a new serviceLister with given api and filters.
func newServiceLister(api *API, filters ...ListServicesFilter) *serviceLister {
	l := &serviceLister{
		api: api,
	}
	for _, filter := range filters {
		filter(l)
	}
	return l
}

// Lists services.
func (l *serviceLister) List() ([]*service.Service, error) {
	var ids []string

	if l.filterActive {
		var err error
		ids, err = l.getActiveServicesIDs()
		if err != nil {
			return nil, err
		}
		if len(ids) == 0 {
			return nil, nil
		}
	}

	var (
		ss  []*service.Service
		err error
	)

	if len(ids) > 0 {
		ss, err = l.api.db.GetByIDs(ids)
	} else {
		ss, err = l.api.db.All()
	}
	if err != nil {
		return nil, err
	}

	var services []*service.Service
	for _, s := range ss {
		s, err = service.FromService(s, service.ContainerOption(l.api.container))
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}

	return services, nil
}

// getActiveServicesIDs returns the services ids for the services that are
// in the starting process or already started.
func (l *serviceLister) getActiveServicesIDs() ([]string, error) {
	var ids []string

	runningServices, err := l.api.container.ListServices("mesg.hash", fmt.Sprintf("mesg.core=%s", l.api.cfg.Core.Name))
	if err != nil {
		return nil, err
	}

	// Make service list unique. One mesg service can have multiple docker service.
	runningServiceIDs := make(map[string]bool)
	for _, service := range runningServices {
		serviceName := service.Spec.Annotations.Labels["mesg.hash"]
		runningServiceIDs[serviceName] = true
	}

	for id := range runningServiceIDs {
		ids = append(ids, id)
	}

	return ids, nil
}
