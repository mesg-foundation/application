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

package event

import (
	"testing"

	"github.com/mesg-foundation/core/service"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	s, _ := service.FromService(&service.Service{
		Name: "TestCreate",
		Events: []*service.Event{
			{Key: "test"},
		},
	})
	var data map[string]interface{}
	exec, err := Create(s, "test", data)
	require.NoError(t, err)
	require.Equal(t, s, exec.Service)
	require.Equal(t, data, exec.Data)
	require.Equal(t, "test", exec.Key)
	require.NotNil(t, exec.CreatedAt)
}

func TestCreateNotPresentEvent(t *testing.T) {
	var (
		serviceName      = "TestCreateNotPresentEvent"
		eventKey         = "test"
		invalidEventName = "testInvalid"
	)
	s, _ := service.FromService(&service.Service{
		Name: serviceName,
		Events: []*service.Event{
			{
				Key: eventKey,
			},
		},
	})
	var data map[string]interface{}
	_, err := Create(s, invalidEventName, data)
	require.Error(t, err)
	notFoundErr, ok := err.(*service.EventNotFoundError)
	require.True(t, ok)
	require.Equal(t, invalidEventName, notFoundErr.EventKey)
	require.Equal(t, serviceName, notFoundErr.ServiceName)
}

func TestCreateInvalidData(t *testing.T) {
	var (
		eventKey    = "test"
		serviceName = "TestCreateInvalidData"
	)
	s, _ := service.FromService(&service.Service{
		Name: serviceName,
		Events: []*service.Event{
			{
				Key: eventKey,
				Data: []*service.Parameter{
					{Key: "xxx"},
				},
			},
		},
	})
	var data map[string]interface{}
	_, err := Create(s, "test", data)
	require.Error(t, err)
	invalidErr, ok := err.(*service.InvalidEventDataError)
	require.True(t, ok)
	require.Equal(t, eventKey, invalidErr.EventKey)
	require.Equal(t, serviceName, invalidErr.ServiceName)
}
