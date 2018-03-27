package service

import (
	"strings"

	"github.com/fsouza/go-dockerclient"
)

func deleteNetwork(namespace string) (err error) {
	cli, err := dockerCli()
	if err != nil {
		return
	}
	network, err := findNetwork(namespace)
	if err != nil {
		return
	}
	return cli.RemoveNetwork(network.ID)
}

func findNetwork(namespace string) (network *docker.Network, err error) {
	cli, err := dockerCli()
	if err != nil {
		return
	}
	networks, err := cli.FilteredListNetworks(docker.NetworkFilterOpts{
		"scope": {"swarm": true},
		"label": {"com.docker.stack.namespace=" + namespace: true},
	})
	if len(networks) > 0 {
		network = &networks[0]
	}
	return
}

func createNetwork(namespace string) (network *docker.Network, err error) {
	cli, err := dockerCli()
	if err != nil {
		return
	}
	network, err = findNetwork(namespace)
	if network != nil || err != nil {
		return
	}
	network, err = cli.CreateNetwork(docker.CreateNetworkOptions{
		Name:   strings.Join([]string{namespace, "Network"}, "-"),
		Driver: "overlay",
		Labels: map[string]string{
			"com.docker.stack.namespace": namespace,
		},
	})
	return
}
