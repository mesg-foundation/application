package container

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/mesg-foundation/core/container/dockertest"
	"github.com/stretchr/testify/require"
)

func TestStartService(t *testing.T) {
	t.Skip("put back when dockertest will be replaced with testify.Mock")
	namespace := []string{"namespace"}
	containerID := "id"
	options := ServiceOptions{
		Image:     "http-server",
		Namespace: namespace,
	}

	dt := dockertest.New()
	c, _ := New(ClientOption(dt.Client()))

	dt.ProvideServiceCreate(types.ServiceCreateResponse{ID: containerID}, nil)

	id, err := c.StartService(options)
	require.NoError(t, err)
	require.Equal(t, containerID, id)

	ls := <-dt.LastServiceCreate()
	require.Equal(t, options.toSwarmServiceSpec(c), ls.Service)
	require.Equal(t, types.ServiceCreateOptions{}, ls.Options)
}

func TestStopService(t *testing.T) {
	t.Skip("put back when dockertest will be replaced with testify.Mock")

	namespace := []string{"namespace"}
	dt := dockertest.New()
	c, _ := New(ClientOption(dt.Client()))
	containerID := "1"

	dt.ProvideContainerList([]types.Container{
		{ID: containerID},
	}, nil)
	dt.ProvideContainerInspect(types.ContainerJSON{
		ContainerJSONBase: &types.ContainerJSONBase{ID: containerID},
	}, nil)

	dt.ProvideContainerList(nil, dockertest.NotFoundErr{})
	dt.ProvideServiceInspectWithRaw(swarm.Service{}, nil, dockertest.NotFoundErr{})

	require.Nil(t, c.StopService(namespace))
	require.Equal(t, c.Namespace(namespace), (<-dt.LastServiceRemove()).ServiceID)

	ls := <-dt.LastContainerStop()
	require.Equal(t, containerID, ls.Container)

	lr := <-dt.LastContainerRemove()
	require.Equal(t, containerID, lr.Container)
	require.Equal(t, types.ContainerRemoveOptions{}, lr.Options)
}

func TestStopNotExistingService(t *testing.T) {
	namespace := []string{"namespace"}

	dt := dockertest.New()
	c, _ := New(ClientOption(dt.Client()))

	dt.ProvideContainerList([]types.Container{}, nil)
	dt.ProvideContainerInspect(types.ContainerJSON{}, dockertest.NotFoundErr{})
	dt.ProvideServiceRemove(dockertest.NotFoundErr{})
	dt.ProvideServiceInspectWithRaw(swarm.Service{}, nil, dockertest.NotFoundErr{})
	dt.ProvideContainerInspect(types.ContainerJSON{}, dockertest.NotFoundErr{})

	require.NoError(t, c.StopService(namespace))
}

func TestFindService(t *testing.T) {
	namespace := []string{"namespace"}
	swarmService := swarm.Service{ID: "1"}

	dt := dockertest.New()
	c, _ := New(ClientOption(dt.Client()))

	dt.ProvideServiceInspectWithRaw(swarmService, nil, nil)

	service, err := c.FindService(namespace)
	require.NoError(t, err)
	require.Equal(t, swarmService.ID, service.ID)

	li := <-dt.LastServiceInspectWithRaw()
	require.Equal(t, c.Namespace(namespace), li.ServiceID)
	require.Equal(t, types.ServiceInspectOptions{}, li.Options)
}

func TestFindServiceNotExisting(t *testing.T) {
	namespace := []string{"namespace"}

	dt := dockertest.New()
	c, _ := New(ClientOption(dt.Client()))

	dt.ProvideServiceInspectWithRaw(swarm.Service{}, nil, dockertest.NotFoundErr{})

	_, err := c.FindService(namespace)
	require.Equal(t, dockertest.NotFoundErr{}, err)

	li := <-dt.LastServiceInspectWithRaw()
	require.Equal(t, c.Namespace(namespace), li.ServiceID)
	require.Equal(t, types.ServiceInspectOptions{}, li.Options)
}

func TestListServices(t *testing.T) {
	namespace := []string{"namespace"}
	namespace1 := []string{"namespace"}
	label := "1"
	dt := dockertest.New()
	c, _ := New(ClientOption(dt.Client()))
	swarmServices := []swarm.Service{
		{Spec: swarm.ServiceSpec{Annotations: swarm.Annotations{Name: c.Namespace(namespace)}}},
		{Spec: swarm.ServiceSpec{Annotations: swarm.Annotations{Name: c.Namespace(namespace1)}}},
	}

	dt.ProvideServiceList(swarmServices, nil)

	services, err := c.ListServices(label)
	require.NoError(t, err)
	require.Equal(t, 2, len(services))
	require.Equal(t, c.Namespace(namespace), services[0].Spec.Name)
	require.Equal(t, c.Namespace(namespace1), services[1].Spec.Name)

	require.Equal(t, types.ServiceListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "label",
			Value: label,
		}),
	}, (<-dt.LastServiceList()).Options)
}

func TestServiceLogs(t *testing.T) {
	namespace := []string{"namespace"}
	data := []byte{1, 2}

	dt := dockertest.New()
	c, _ := New(ClientOption(dt.Client()))

	dt.ProvideServiceLogs(ioutil.NopCloser(bytes.NewReader(data)), nil)

	reader, err := c.ServiceLogs(namespace)
	require.NoError(t, err)
	defer reader.Close()

	bytes, err := ioutil.ReadAll(reader)
	require.NoError(t, err)
	require.Equal(t, data, bytes)

	ll := <-dt.LastServiceLogs()
	require.Equal(t, c.Namespace(namespace), ll.ServiceID)
	require.Equal(t, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: false,
		Follow:     true,
	}, ll.Options)
}
