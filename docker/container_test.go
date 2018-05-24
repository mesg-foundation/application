package docker

import (
	"testing"
	"time"

	"github.com/stvp/assert"
)

func TestFindContainerNotExisting(t *testing.T) {
	namespace := []string{"TestFindContainerNotExisting"}
	container, err := FindContainer(namespace)
	assert.Nil(t, err)
	assert.Nil(t, container)
}

func TestFindContainer(t *testing.T) {
	namespace := []string{"TestFindContainer"}
	startTestService(namespace)
	defer StopService(namespace)
	<-WaitContainerStatus(namespace, RUNNING, time.Minute)
	container, err := FindContainer(namespace)
	assert.Nil(t, err)
	assert.NotNil(t, container)
}

func TestFindContainerStopped(t *testing.T) {
	namespace := []string{"TestFindContainerStopped"}
	startTestService(namespace)
	StopService(namespace)
	container, err := FindContainer(namespace)
	assert.Nil(t, err)
	assert.Nil(t, container)
}

func TestContainerStatusNeverStarted(t *testing.T) {
	namespace := []string{"TestContainerStatusNeverStarted"}
	status, err := ContainerStatus(namespace)
	assert.Nil(t, err)
	assert.Equal(t, status, STOPPED)
}

func TestContainerStatusRunning(t *testing.T) {
	namespace := []string{"TestContainerStatusRunning"}
	startTestService(namespace)
	defer StopService(namespace)
	<-WaitContainerStatus(namespace, RUNNING, time.Minute)
	status, err := ContainerStatus(namespace)
	assert.Nil(t, err)
	assert.Equal(t, status, RUNNING)
}

func TestContainerStatusStopped(t *testing.T) {
	namespace := []string{"TestContainerStatusStopped"}
	startTestService(namespace)
	<-WaitContainerStatus(namespace, RUNNING, time.Minute)

	StopService(namespace)
	<-WaitContainerStatus(namespace, STOPPED, time.Minute)

	status, err := ContainerStatus(namespace)
	assert.Nil(t, err)
	assert.Equal(t, status, STOPPED)
}

func TestWaitForContainerRunning(t *testing.T) {
	namespace := []string{"TestWaitForContainerRunning"}
	startTestService(namespace)
	defer StopService(namespace)

	wait := WaitContainerStatus(namespace, RUNNING, time.Minute)
	err := <-wait
	assert.Nil(t, err)
}

func TestWaitForContainerStopped(t *testing.T) {
	namespace := []string{"TestWaitForContainerStopped"}
	startTestService(namespace)
	<-WaitContainerStatus(namespace, RUNNING, time.Minute)

	StopService(namespace)
	wait := WaitContainerStatus(namespace, STOPPED, time.Minute)
	err := <-wait
	assert.Nil(t, err)
}
