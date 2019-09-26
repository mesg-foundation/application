package servicesdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/docker/docker/pkg/archive"
	"github.com/mesg-foundation/engine/container"
	"github.com/mesg-foundation/engine/database"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/hash/dirhash"
	"github.com/mesg-foundation/engine/protobuf/api"
	ownershipsdk "github.com/mesg-foundation/engine/sdk/ownership"
	"github.com/mesg-foundation/engine/service"
	"github.com/mesg-foundation/engine/service/validator"
	"github.com/mr-tron/base58"
)

func create(container container.Container, db *database.ServiceDB, req *api.CreateServiceRequest, owner cosmostypes.AccAddress, ownerships *ownershipsdk.Backend, request cosmostypes.Request) (*service.Service, error) {
	// download and untar service context into path.
	path, err := ioutil.TempDir("", "mesg")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(path)

	resp, err := http.Get("http://ipfs.app.mesg.com:8080/ipfs/" + req.Source)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("service's source code is not reachable")
	}
	defer resp.Body.Close()

	if err := archive.Untar(resp.Body, path, nil); err != nil {
		return nil, err
	}

	// create service
	srv := &service.Service{
		Sid:           req.Sid,
		Name:          req.Name,
		Description:   req.Description,
		Configuration: req.Configuration,
		Tasks:         req.Tasks,
		Events:        req.Events,
		Dependencies:  req.Dependencies,
		Repository:    req.Repository,
		Source:        req.Source,
	}

	// calculate and apply hash to service.
	dh := dirhash.New(path)
	srv.Hash, err = dh.Sum(hash.Dump(srv))
	if err != nil {
		return nil, err
	}

	// check if service already exists.
	if _, err := db.Get(srv.Hash); err == nil {
		return nil, &AlreadyExistsError{Hash: srv.Hash}
	}

	// build service's Docker image.
	_, err = container.Build(path)
	if err != nil {
		return nil, err
	}
	// TODO: the following test should be moved in New function
	if srv.Sid == "" {
		// make sure that sid doesn't have the same length with id.
		srv.Sid = "_" + base58.Encode(srv.Hash) // TODO: use string method after change type to hash.Hash
	}

	if err := validator.ValidateService(srv); err != nil {
		return nil, err
	}

	if _, err := ownerships.CreateServiceOwnership(request, srv.Hash, owner); err != nil {
		return nil, err
	}

	return srv, db.Save(srv)
}
