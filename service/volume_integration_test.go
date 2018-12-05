// +build integration

package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntegrationDeleteVolumes(t *testing.T) {
	// TODO: the following test doesn't work on CircleCI because we don't use "machine"
	// https://support.circleci.com/hc/en-us/articles/360007324514-How-can-I-mount-volumes-to-docker-containers-
	t.Skip("doesn't work on CircleCI because we don't use machine")
	var (
		dependencyKey1 = "1"
		dependencyKey2 = "2"
		volumeA        = "/a"
		volumeB        = "/b"
		s, _           = FromService(&Service{
			Name: "TestIntegrationDeleteVolumes",
			Dependencies: []*Dependency{
				{
					Key:     dependencyKey1,
					Image:   "http-server",
					Volumes: []string{volumeA, volumeB},
				},
				{
					Key:         dependencyKey2,
					Image:       "http-server",
					VolumesFrom: []string{dependencyKey1},
				},
			},
		}, ContainerOption(newIntegrationContainer(t)))
	)
	_, err := s.Start()
	require.NoError(t, err)
	err = s.Stop()
	require.NoError(t, err)
	err = s.DeleteVolumes()
	require.NoError(t, err)
}
