package pod

import "log"

func (pod *Pod) CreatePod() {

	pauseId := createPauseContainer(pod.Metadata.Name)

	for _, container := range pod.Spec.Containers {
		_ = createContainerInPod(pod.Metadata.Name, container, pauseId)
		log.Printf("Created container %s in pod %s", container.Name, pod.Metadata.Name)
	}
}
