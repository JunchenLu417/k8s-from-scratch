package docker_go_client

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

func MakeBasics(imageName string, cmd []string) (*container.Config, *container.HostConfig, *network.NetworkingConfig, error) {

	// Basic container config
	containerConfig := &container.Config{
		Image: imageName,
		Cmd:   cmd,
	}

	// Basic host config
	hostConfig := &container.HostConfig{}

	// Basic networking config
	networkingConfig := &network.NetworkingConfig{}

	return containerConfig, hostConfig, networkingConfig, nil
}
