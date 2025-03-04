package docker_go_client

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"os"
	"strings"
	"text/tabwriter"
)

func getContainerList(cli *client.Client, getAll bool) ([]types.Container, error) {
	opts := container.ListOptions{}
	if getAll {
		opts.All = true
	} // otherwise, get running containers
	return cli.ContainerList(context.Background(), opts)
}

func GetContainerId(containerName string) (string, error) {
	cli, err := getNewClient()
	if err != nil {
		return "", err
	}
	list, err := getContainerList(cli, true)
	if err != nil {
		return "", err
	}
	for _, c := range list {
		if strings.TrimPrefix(c.Names[0], "/") == containerName {
			return c.ID, nil
		}
	}
	return "", fmt.Errorf("container not found")
}

func PrettyPrint() error {
	cli, err := getNewClient()
	if err != nil {
		return err
	}
	list, err := getContainerList(cli, true)
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "CONTAINER ID\tIMAGE\tSTATUS\tNAMES")

	for _, c := range list {
		_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			c.ID[:12], c.Image, c.Status, strings.TrimPrefix(c.Names[0], "/"))
	}

	_ = w.Flush()
	return nil
}
