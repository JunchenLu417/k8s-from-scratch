package pod_test

import (
	"github.com/JunchenLu417/k8s-from-scratch/pkg/apis"
	mypod "github.com/JunchenLu417/k8s-from-scratch/pkg/apis/pod"
	"testing"
)

func TestCreatePod(t *testing.T) {

	pod, err := apis.DecodeApiFromYaml(
		"/home/ubuntu/k8s-from-scratch/pkg/apis/pod/testing/yamls/localhost_test.yaml",
	)
	if err != nil {
		t.Errorf("Decode failed: %v", err)
	}

	podToCreate, _ := pod.(*mypod.Pod)
	podToCreate.CreatePod()
}
