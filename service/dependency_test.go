package service

import (
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types/swarm"
	"github.com/stvp/assert"
)

// TODO: this shouldn't be necessary
func TestMain(m *testing.M) {
	// This is to initialize the swarm if needed
	dockerCli()
	time.Sleep(1 * time.Second)
}

func TestExtractPortEmpty(t *testing.T) {
	ports := extractPorts(&Dependency{})
	assert.Equal(t, len(ports), 0)
}

func TestExtractPorts(t *testing.T) {
	ports := extractPorts(&Dependency{
		Ports: []string{
			"80",
			"3000:8080",
		},
	})
	assert.Equal(t, len(ports), 2)
	assert.Equal(t, ports[0].PublishMode, swarm.PortConfigPublishModeIngress)
	assert.Equal(t, ports[0].Protocol, swarm.PortConfigProtocolTCP)
	assert.Equal(t, ports[0].TargetPort, uint32(80))
	assert.Equal(t, ports[0].PublishedPort, uint32(80))
	assert.Equal(t, ports[1].TargetPort, uint32(8080))
	assert.Equal(t, ports[1].PublishedPort, uint32(3000))
}

func TestGetDockerService(t *testing.T) {
	namespace := strings.Join([]string{NAMESPACE, "TestGetDockerService"}, "_")
	name := "test"
	dependency := Dependency{Image: "nginx"}
	network, err := createNetwork(namespace)
	dependency.Start(dependencyDetails{
		namespace:      namespace,
		dependencyName: name,
		serviceName:    "TestGetDockerService",
	}, network)
	res, err := getDockerService(namespace, name)
	assert.Nil(t, err)
	assert.NotEqual(t, res.ID, "")
	res, err = getDockerService(namespace, "textx")
	assert.Nil(t, err)
	assert.Equal(t, res.ID, "")
	dependency.Stop(namespace, name)
	deleteNetwork(namespace)
}

func TestDockerServiceMatch(t *testing.T) {
	namespace := strings.Join([]string{NAMESPACE, "TestDockerServiceMatch"}, "_")
	dockerServices := []swarm.Service{
		swarm.Service{
			Spec: swarm.ServiceSpec{
				Annotations: swarm.Annotations{
					Name: strings.Join([]string{namespace, "test1"}, "_"),
				},
			},
		},
		swarm.Service{
			Spec: swarm.ServiceSpec{
				Annotations: swarm.Annotations{
					Name: strings.Join([]string{namespace, "test2"}, "_"),
				},
			},
		},
	}
	res1 := dockerServiceMatch(dockerServices, namespace, "test")
	assert.Equal(t, res1, swarm.Service{})
	res2 := dockerServiceMatch(dockerServices, namespace, "test1")
	assert.Equal(t, res2, dockerServices[0])
	res3 := dockerServiceMatch(dockerServices, namespace, "test2")
	assert.Equal(t, res3, dockerServices[1])
}
