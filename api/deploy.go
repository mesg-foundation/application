package api

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/service/importer"
	"github.com/mesg-foundation/core/x/xdocker/xarchive"
	"github.com/mesg-foundation/core/x/xgit"
	uuid "github.com/satori/go.uuid"
)

// DeployServiceOption is a configuration func for deploying services.
type DeployServiceOption func(*serviceDeployer)

// DeployServiceStatusOption receives chan statuses to send deploy statuses.
func DeployServiceStatusOption(statuses chan DeployStatus) DeployServiceOption {
	return func(deployer *serviceDeployer) {
		deployer.statuses = statuses
	}
}

// DeployService deploys a service from a gzipped tarball.
func (a *API) DeployService(r io.Reader, options ...DeployServiceOption) (*service.Service,
	*importer.ValidationError, error) {
	return newServiceDeployer(a, options...).FromGzippedTar(r)
}

// DeployServiceFromURL deploys a service living at a Git host.
// Supported URL types:
// - https://github.com/mesg-foundation/service-ethereum
// - https://github.com/mesg-foundation/service-ethereum#branchName
func (a *API) DeployServiceFromURL(url string, options ...DeployServiceOption) (*service.Service,
	*importer.ValidationError, error) {
	return newServiceDeployer(a, options...).FromGitURL(url)
}

// serviceDeployer provides functionalities to deploy a MESG service.
type serviceDeployer struct {
	// statuses receives status messages produced during deployment.
	statuses chan DeployStatus

	api *API
}

// StatusType indicates the type of status message.
type StatusType int

const (
	_ StatusType = iota // skip zero value.

	// Running indicates that status message belongs to a continuous state.
	Running

	// DonePositive indicates that status message belongs to a positive noncontinuous state.
	DonePositive

	// DoneNegative indicates that status message belongs to a negative noncontinuous state.
	DoneNegative
)

// DeployStatus represents the deployment status.
type DeployStatus struct {
	Message string
	Type    StatusType
}

// newServiceDeployer creates a new serviceDeployer with given api and options.
func newServiceDeployer(api *API, options ...DeployServiceOption) *serviceDeployer {
	d := &serviceDeployer{
		api: api,
	}
	for _, option := range options {
		option(d)
	}
	return d
}

// FromGitURL deploys a service hosted at a Git url.
func (d *serviceDeployer) FromGitURL(url string) (*service.Service, *importer.ValidationError, error) {
	defer d.closeStatus()
	d.sendStatus("Downloading service...", Running)
	path, err := d.createTempDir()
	if err != nil {
		return nil, nil, err
	}
	defer os.RemoveAll(path)
	if err := xgit.Clone(url, path); err != nil {
		return nil, nil, err
	}

	// XXX: remove .git folder from repo.
	// It makes docker build iamge id same between repo clones.
	if err := os.RemoveAll(filepath.Join(path, ".git")); err != nil {
		return nil, nil, err
	}

	d.sendStatus("Service downloaded with success", DonePositive)
	r, err := xarchive.GzippedTar(path)
	if err != nil {
		return nil, nil, err
	}
	return d.deploy(r)
}

// FromGzippedTar deploys a service from a gzipped tarball.
func (d *serviceDeployer) FromGzippedTar(r io.Reader) (*service.Service, *importer.ValidationError, error) {
	defer d.closeStatus()
	return d.deploy(r)
}

// deploy deploys a service in path.
func (d *serviceDeployer) deploy(r io.Reader) (*service.Service, *importer.ValidationError, error) {
	var (
		statuses = make(chan service.DeployStatus)
		wg       sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		d.forwardDeployStatuses(statuses)
	}()

	s, err := service.New(r,
		service.ContainerOption(d.api.container),
		service.DeployStatusOption(statuses),
	)
	wg.Wait()

	validationErr, err := d.assertValidationError(err)
	if err != nil {
		return nil, nil, err
	}
	if validationErr != nil {
		return nil, validationErr, nil
	}

	return s, nil, d.api.db.Save(s)
}

func (d *serviceDeployer) createTempDir() (path string, err error) {
	return ioutil.TempDir("", "mesg-"+uuid.NewV4().String())
}

// sendStatus sends a status message.
func (d *serviceDeployer) sendStatus(message string, typ StatusType) {
	if d.statuses != nil {
		d.statuses <- DeployStatus{
			Message: message,
			Type:    typ,
		}
	}
}

// closeStatus closes statuses chan.
func (d *serviceDeployer) closeStatus() {
	if d.statuses != nil {
		close(d.statuses)
	}
}

// forwardStatuses forwards status messages.
func (d *serviceDeployer) forwardDeployStatuses(statuses chan service.DeployStatus) {
	for status := range statuses {
		var t StatusType
		switch status.Type {
		case service.DRunning:
			t = Running
		case service.DDonePositive:
			t = DonePositive
		case service.DDoneNegative:
			t = DoneNegative
		}
		d.sendStatus(status.Message, t)
	}
}

func (d *serviceDeployer) assertValidationError(err error) (*importer.ValidationError, error) {
	if err == nil {
		return nil, nil
	}
	if validationError, ok := err.(*importer.ValidationError); ok {
		return validationError, nil
	}
	return nil, err
}
