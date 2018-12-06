// +build integration

package container

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func startTestService(name []string) (string, error) {
	c, err := New()
	if err != nil {
		return "", err
	}
	return c.StartService(ServiceOptions{
		Image:     "http-server",
		Namespace: name,
	})
}

func TestIntegrationStartService(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestStartService"}
	serviceID, err := startTestService(namespace)
	defer c.StopService(namespace)
	require.NoError(t, err)
	require.NotEqual(t, "", serviceID)
}

func TestIntegrationStartServiceTwice(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestStartServiceTwice"}
	id1, err := startTestService(namespace)
	require.NoError(t, err)
	defer c.StopService(namespace)
	id2, err := startTestService(namespace)
	require.NoError(t, err)
	require.Equal(t, id1, id2)
}

func TestIntegrationStopService(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestStopService"}
	startTestService(namespace)
	err = c.StopService(namespace)
	require.NoError(t, err)
}

func TestIntegrationStopNotExistingService(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestStopNotExistingService"}
	err = c.StopService(namespace)
	require.NoError(t, err)
}

func TestIntegrationStatusNeverStarted(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestStatusNeverStarted"}
	status, err := c.Status(namespace)
	require.NoError(t, err)
	require.NotEqual(t, RUNNING, status)
	require.Equal(t, STOPPED, status)
}

func TestIntegrationStatusRunning(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestStatusRunning"}
	startTestService(namespace)
	defer c.StopService(namespace)
	status, err := c.Status(namespace)
	require.NoError(t, err)
	require.Equal(t, status, RUNNING)
	require.NotEqual(t, status, STOPPED)
}

func TestIntegrationStatusStopped(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestStatusStopped"}
	startTestService(namespace)
	c.StopService(namespace)
	status, err := c.Status(namespace)
	require.NoError(t, err)
	require.Equal(t, status, STOPPED)
	require.NotEqual(t, status, RUNNING)
}

func TestIntegrationFindServiceNotExisting(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	_, err = c.FindService([]string{"TestFindServiceNotExisting"})
	require.Error(t, err)
}

func TestIntegrationFindService(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestFindService"}
	startTestService(namespace)
	defer c.StopService(namespace)
	service, err := c.FindService(namespace)
	require.NoError(t, err)
	require.NotEqual(t, "", service.ID)
}

func TestIntegrationFindServiceCloseName(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestFindServiceCloseName", "name"}
	namespace1 := []string{"TestFindServiceCloseName", "name2"}
	startTestService(namespace)
	defer c.StopService(namespace)
	startTestService(namespace1)
	defer c.StopService(namespace1)
	service, err := c.FindService(namespace)
	require.NoError(t, err)
	require.NotEqual(t, "", service.ID)
}

func TestIntegrationFindServiceStopped(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestFindServiceStopped"}
	startTestService(namespace)
	c.StopService(namespace)
	_, err = c.FindService(namespace)
	require.Error(t, err)
}

func TestIntegrationListServices(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	c.StartService(ServiceOptions{
		Image:     "http-server",
		Namespace: []string{"TestListServices"},
		Labels: map[string]string{
			"label_name": "value_1",
		},
	})
	c.StartService(ServiceOptions{
		Image:     "http-server",
		Namespace: []string{"TestListServiceswithValue2"},
		Labels: map[string]string{
			"label_name_2": "value_2",
		},
	})
	defer c.StopService([]string{"TestListServices"})
	defer c.StopService([]string{"TestListServiceswithValue2"})
	services, err := c.ListServices("label_name")
	require.NoError(t, err)
	require.Equal(t, 1, len(services))
	require.Equal(t, c.Namespace([]string{"TestListServices"}), services[0].Spec.Name)
}

func TestIntegrationServiceLogs(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	namespace := []string{"TestServiceLogs"}
	startTestService(namespace)
	defer c.StopService(namespace)
	reader, err := c.ServiceLogs(namespace)
	require.NoError(t, err)
	require.NotNil(t, reader)
}
