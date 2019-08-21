package servicesdk

import (
	"context"
	"fmt"

	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/service"
	"github.com/mesg-foundation/engine/database"
)

type Service interface {
	Create(srv *service.Service) (*service.Service, error)
	Delete(hash hash.Hash) error
	Get(hash hash.Hash) (*service.Service, error)
	List() ([]*service.Service, error)
}

type KeeperFactory func(context.Context) (*database.ServiceKeeper, error)

// AlreadyExistsError is an not found error.
type AlreadyExistsError struct {
	Hash hash.Hash
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("service %q already exists", e.Hash.String())
}
