package api

import (
	"github.com/mesg-foundation/core/service"
)

// DeleteService stops and deletes service serviceID.
func (a *API) DeleteService(serviceID string) error {
	return newServiceDeleter(a).Delete(serviceID)
}

// serviceDeleter provides functionalities to delete a MESG service.
type serviceDeleter struct {
	api *API
}

// newServiceDeleter creates a new serviceDeleter with given.
func newServiceDeleter(api *API) *serviceDeleter {
	return &serviceDeleter{
		api: api,
	}
}

// Delete stops and deletes service serviceID.
func (d *serviceDeleter) Delete(serviceID string) error {
	s, err := d.api.db.Get(serviceID)
	if err != nil {
		return err
	}
	s, err = service.FromService(s, service.ContainerOption(d.api.container))
	if err != nil {
		return err
	}
	if err := s.Stop(); err != nil {
		return err
	}
	return d.api.db.Delete(serviceID)
}
