package container

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceOptionNamespace(t *testing.T) {
	namespace := []string{"name1", "name2"}
	options := &ServiceOptions{
		Namespace: namespace,
	}
	c, _ := New()
	expectedNamespace := c.Namespace(namespace)
	service := options.toSwarmServiceSpec(c)
	require.Equal(t, expectedNamespace, service.Annotations.Name)
	require.Equal(t, expectedNamespace, service.Annotations.Labels["com.docker.stack.namespace"])
	require.Equal(t, expectedNamespace, service.TaskTemplate.ContainerSpec.Labels["com.docker.stack.namespace"])
}

func TestServiceOptionImage(t *testing.T) {
	image := "http-server"
	options := &ServiceOptions{
		Image: image,
	}
	c, _ := New()
	service := options.toSwarmServiceSpec(c)
	require.Equal(t, image, service.Annotations.Labels["com.docker.stack.image"])
	require.Equal(t, image, service.TaskTemplate.ContainerSpec.Image)
}

func TestServiceOptionMergeLabels(t *testing.T) {
	l1 := map[string]string{
		"label1": "foo",
		"label2": "bar",
	}
	l2 := map[string]string{
		"label2": "foo",
		"label3": "foo",
		"label4": "bar",
	}
	labels := mergeLabels(l1, l2)
	require.Equal(t, "foo", labels["label1"])
	require.Equal(t, "foo", labels["label2"])
	require.Equal(t, "foo", labels["label3"])
	require.Equal(t, "bar", labels["label4"])
}

func TestServiceOptionLabels(t *testing.T) {
	options := &ServiceOptions{
		Labels: map[string]string{
			"label1": "foo",
			"label2": "bar",
		},
	}
	c, _ := New()
	service := options.toSwarmServiceSpec(c)
	require.Equal(t, "foo", service.Annotations.Labels["label1"])
	require.Equal(t, "bar", service.Annotations.Labels["label2"])
}

func TestServiceOptionPorts(t *testing.T) {
	options := &ServiceOptions{
		Ports: []Port{
			{
				Published: 50503,
				Target:    50501,
			},
			{
				Published: 30503,
				Target:    30501,
			},
		},
	}
	ports := options.swarmPorts()
	require.Equal(t, 2, len(ports))
	require.Equal(t, uint32(50503), ports[0].PublishedPort)
	require.Equal(t, uint32(50501), ports[0].TargetPort)
	require.Equal(t, uint32(30503), ports[1].PublishedPort)
	require.Equal(t, uint32(30501), ports[1].TargetPort)
}

func TestServiceOptionMounts(t *testing.T) {
	options := &ServiceOptions{
		Mounts: []Mount{
			{
				Source: "source/file",
				Target: "target/file",
			},
		},
	}
	mounts := options.swarmMounts(true)
	require.Equal(t, 1, len(mounts))
	require.Equal(t, "source/file", mounts[0].Source)
	require.Equal(t, "target/file", mounts[0].Target)
}

func TestServiceOptionEnv(t *testing.T) {
	options := &ServiceOptions{
		Env: []string{"env1", "env2"},
	}
	c, _ := New()
	service := options.toSwarmServiceSpec(c)
	env := service.TaskTemplate.ContainerSpec.Env
	require.Equal(t, 2, len(env))
	require.Equal(t, "env1", env[0])
	require.Equal(t, "env2", env[1])
}

func TestServiceOptionNetworks(t *testing.T) {
	options := &ServiceOptions{
		Networks: []Network{
			{ID: "network1"},
			{ID: "network2", Alias: "test"},
		},
	}
	networks := options.swarmNetworks()
	require.Equal(t, 2, len(networks))
	require.Equal(t, "network1", networks[0].Target)
	require.Equal(t, 0, len(networks[0].Aliases))
	require.Equal(t, "network2", networks[1].Target)
	require.Equal(t, 1, len(networks[1].Aliases))
	require.Equal(t, "test", networks[1].Aliases[0])
}

func contains(list []string, item string) bool {
	for _, itemInList := range list {
		if itemInList == item {
			return true
		}
	}
	return false
}

func TestMapToEnv(t *testing.T) {
	env := MapToEnv(map[string]string{
		"first":  "first_value",
		"second": "second_value",
	})
	require.True(t, contains(env, "first=first_value"))
	require.True(t, contains(env, "second=second_value"))
}
