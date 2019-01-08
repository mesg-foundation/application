// Copyright 2018 MESG Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"io"
	"sync"

	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/daemon"
	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/x/xerrors"
)

// CoreProvider is a struct that provides all methods required by core command.
type CoreProvider struct {
	d      daemon.Daemon
	client coreapi.CoreClient
}

// NewCoreProvider creates new CoreProvider.
func NewCoreProvider(client coreapi.CoreClient, d daemon.Daemon) *CoreProvider {
	return &CoreProvider{
		client: client,
		d:      d,
	}
}

// Start starts core daemon.
func (p *CoreProvider) Start() error {
	return p.d.Start()
}

// Stop stops core daemon and all running services.
func (p *CoreProvider) Stop() error {
	ids, err := service.ListRunning()
	if err != nil {
		return err
	}

	var (
		idsLen = len(ids)
		errC   = make(chan error, idsLen)
		wg     sync.WaitGroup
	)

	wg.Add(idsLen)
	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			_, err := p.client.StopService(context.Background(), &coreapi.StopServiceRequest{
				ServiceID: id,
			})
			if err != nil {
				errC <- err
			}
		}(id)
	}

	wg.Wait()
	close(errC)

	var errs xerrors.Errors
	for err := range errC {
		errs = append(errs, err)
	}

	if err := p.d.Stop(); err != nil {
		errs = append(errs, err)
	}

	return errs.ErrorOrNil()
}

// Status returns daemon status.
func (p *CoreProvider) Status() (container.StatusType, error) {
	return p.d.Status()
}

// Logs returns daemon logs reader.
func (p *CoreProvider) Logs() (io.ReadCloser, error) {
	return p.d.Logs()
}
