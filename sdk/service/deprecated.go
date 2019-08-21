package servicesdk

import (
	"github.com/mesg-foundation/engine/container"
	"github.com/mesg-foundation/engine/database"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/service"
)

type deprecated struct {
	container container.Container
	keeper    *database.ServiceKeeper
}

func NewDeprecated(c container.Container, keeper *database.ServiceKeeper) Service {
	return &deprecated{
		container: c,
		keeper:    keeper,
	}
}

// Create creates a new service from definition.
func (s *deprecated) Create(srv *service.Service) (*service.Service, error) {
	return create(s.container, s.keeper, srv)
}

// Delete deletes the service by hash.
func (s *deprecated) Delete(hash hash.Hash) error {
	return s.keeper.Delete(hash)
}

// Get returns the service that matches given hash.
func (s *deprecated) Get(hash hash.Hash) (*service.Service, error) {
	return s.keeper.Get(hash)
}

// List returns all services.
func (s *deprecated) List() ([]*service.Service, error) {
	return s.keeper.All()
}
