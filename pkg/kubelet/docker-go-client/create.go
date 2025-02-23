package docker_go_client

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
	"io"
	"log"
)

func getNewClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.WithVersion("1.47"))
}

func pullImage(client *client.Client, name string) error {
	// Check if the image exists locally
	_, _, err := client.ImageInspectWithRaw(context.Background(), name)
	if err == nil {
		log.Printf("Image %s already exists locally, skipping pull.", name)
		return nil
	}

	// If the image is not found, pull it
	if errdefs.IsNotFound(err) {
		log.Printf("Image %s not found locally, pulling...", name)
		pull, err := client.ImagePull(context.Background(), name, image.PullOptions{})
		if err != nil {
			return err
		}
		defer func(pull io.ReadCloser) {
			err := pull.Close()
			if err != nil {
				log.Printf("failed to close image pull response: %v", err)
			}
		}(pull)

		// Read the output stream to ensure the pull is completed before returning
		_, err = io.Copy(io.Discard, pull)
		if err != nil {
			log.Printf("failed to read image pull response: %v", err)
			return err
		}
	} else {
		return err
	}
	return nil
}

func createContainer(client *client.Client, containerConfig *container.Config, hostConfig *container.HostConfig,
	networkingConfig *network.NetworkingConfig, name string) (string, error) {

	create, err := client.ContainerCreate(context.Background(), containerConfig, hostConfig, networkingConfig, nil, name)
	if err != nil {
		return "", fmt.Errorf("failed to create container: %w", err)
	}
	log.Printf("Container %s created with ID: %s", name, create.ID)
	return create.ID, nil
}

func startContainer(client *client.Client, id string) error {
	err := client.ContainerStart(context.Background(), id, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}
	return nil
}

// RunContainer should be the only exported symbol (visible and invoked outside the package)
func RunContainer(containerName string, containerConfig *container.Config,
	hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig) error {

	cli, err := getNewClient()
	if err != nil {
		return err
	}
	err = pullImage(cli, containerConfig.Image)
	if err != nil {
		return err
	}
	containerID, err := createContainer(cli, containerConfig, hostConfig, networkingConfig, containerName)
	if err != nil {
		return err
	}
	err = startContainer(cli, containerID)
	if err != nil {
		return err
	}
	return nil
}
