package docker_go_client

import "github.com/docker/docker/client"

func GetNewClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.WithVersion("1.47"))
}
