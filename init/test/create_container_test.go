package init_test

import (
	dockerclient "github.com/JunchenLu417/k8s-from-scratch/pkg/kubelet/docker-go-client"
	"github.com/docker/docker/client"
	"testing"
)

func TestCreateContainer(t *testing.T) {
	cli, err := dockerclient.GetNewClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
	}
	// defer: execute anonymous func just before the surrounding func returns
	defer func(cli *client.Client) {
		err := cli.Close()
		if err != nil {
			t.Fatalf("Failed to close Docker client: %v", err)
		}
	}(cli)

	//
}
