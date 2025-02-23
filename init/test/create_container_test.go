package init_test

import (
	dockerClient "github.com/JunchenLu417/k8s-from-scratch/pkg/kubelet/docker-go-client"
	"testing"
)

func TestCreateContainer(t *testing.T) {

	imageName := "busybox"
	containerName := "testcontainer"
	cmd := []string{"sh", "-c", "tail -f /dev/null"} // our long-running command

	containerConfig, hostConfig, networkingConfig, _ := dockerClient.MakeBasics(imageName, cmd)
	err := dockerClient.RunContainer(containerName, containerConfig, hostConfig, networkingConfig)
	if err != nil {
		t.Fatalf("RunContainer failed: %v", err)
	}
}
