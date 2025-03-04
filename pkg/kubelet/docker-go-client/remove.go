package docker_go_client

import (
	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
)

func DeleteContainer(containerId string, remove bool) error {
	cli, err := getNewClient()
	if err != nil {
		return err
	}
	err = cli.ContainerStop(context.Background(), containerId, container.StopOptions{})
	if err != nil {
		return err
	}
	if remove { // remove the container permanently
		err = cli.ContainerRemove(context.Background(), containerId, container.RemoveOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
