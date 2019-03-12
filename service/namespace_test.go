package service

import (
	"crypto/sha1"
	"encoding/hex"
	"testing"

	"github.com/mesg-foundation/core/utils/hash"
	"github.com/stretchr/testify/require"
)

func TestServiceNamespace(t *testing.T) {
	service, _ := FromService(&Service{Name: "TestServiceNamespace"})
	namespace := service.namespace()
	h, err := hex.DecodeString(service.Hash)
	require.NoError(t, err)
	sum := sha1.Sum(h)
	require.Equal(t, namespace, []string{hex.EncodeToString(sum[:])})
}

func TestDependencyNamespace(t *testing.T) {
	service, _ := FromService(&Service{
		Name: "TestDependencyNamespace",
		Dependencies: []*Dependency{
			{
				Key:   "test",
				Image: "http-server",
			},
		},
	})
	dep := service.Dependencies[0]
	h, err := hex.DecodeString(service.Hash)
	require.NoError(t, err)
	sum := sha1.Sum(h)
	require.Equal(t, dep.namespace(), []string{hex.EncodeToString(sum[:]), "test"})
}

func TestEventSubscriptionChannel(t *testing.T) {
	service, _ := FromService(&Service{Name: "TestEventSubscriptionChannel"})
	require.Equal(t, service.EventSubscriptionChannel(), hash.Calculate(append(
		service.namespace(),
		eventChannel,
	)))
}

func TestTaskSubscriptionChannel(t *testing.T) {
	service, _ := FromService(&Service{Name: "TaskSubscriptionChannel"})
	require.Equal(t, service.TaskSubscriptionChannel(), hash.Calculate(append(
		service.namespace(),
		taskChannel,
	)))
}

func TestResultSubscriptionChannel(t *testing.T) {
	service, _ := FromService(&Service{Name: "ResultSubscriptionChannel"})
	require.Equal(t, service.ResultSubscriptionChannel(), hash.Calculate(append(
		service.namespace(),
		resultChannel,
	)))
}
