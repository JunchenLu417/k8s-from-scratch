package pod

import (
	"fmt"
	dockerClient "github.com/JunchenLu417/k8s-from-scratch/pkg/kubelet/docker-go-client"
	"github.com/docker/docker/api/types/container"
)

func createPauseContainer(podName string) string {

	imageName := "k8s.gcr.io/pause:3.8"
	containerName := fmt.Sprintf("%s-%s", podName, "pause")
	var cmd []string

	containerConfig, hostConfig, networkingConfig, _ := dockerClient.MakeBasics(imageName, cmd)
	hostConfig.IpcMode = "shareable"
	_ = dockerClient.RunContainer(containerName, containerConfig, hostConfig, networkingConfig)

	containerId, _ := dockerClient.GetContainerId(containerName)
	return containerId
}

func createContainerInPod(podName string, c *Container, pauseId string) string {

	containerName := fmt.Sprintf("%s-%s", podName, c.Name)

	containerConfig, hostConfig, networkingConfig, _ := dockerClient.MakeBasics(c.Image, c.Command)
	// join the container with pause container
	hostConfig.NetworkMode = container.NetworkMode("container:" + pauseId)
	hostConfig.PidMode = container.PidMode("container:" + pauseId)
	hostConfig.IpcMode = container.IpcMode("container:" + pauseId)

	_ = dockerClient.RunContainer(containerName, containerConfig, hostConfig, networkingConfig)

	containerId, _ := dockerClient.GetContainerId(containerName)
	return containerId
}
