package init_test

import (
	dockerClient "github.com/JunchenLu417/k8s-from-scratch/pkg/kubelet/docker-go-client"
	"testing"
)

func TestRemoveContainer(t *testing.T) {

	err := dockerClient.PrettyPrint()
	if err != nil {
		t.Error(err)
	}

	containerName := "testcontainer"
	containerId, err := dockerClient.GetContainerId(containerName)
	if err != nil {
		t.Error(err)
	}
	err = dockerClient.DeleteContainer(containerId, true)
	if err != nil {
		t.Error(err)
	}

	err = dockerClient.PrettyPrint()
	if err != nil {
		t.Error(err)
	}
}
