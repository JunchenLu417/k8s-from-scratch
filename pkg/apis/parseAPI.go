package apis

import (
	"fmt"
	mypod "github.com/JunchenLu417/k8s-from-scratch/pkg/apis/pod"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type GenericAPI struct {
	Kind string `yaml:"kind"`
}

func DecodeApiFromYaml(filename string) (interface{}, error) {

	yamlBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	var api GenericAPI
	err = yaml.Unmarshal(yamlBytes, &api)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal into generic object: %w", err)
	}

	switch api.Kind {

	case "Pod":
		var pod mypod.Pod
		if err := yaml.Unmarshal(yamlBytes, &pod); err != nil {
			return nil, fmt.Errorf("failed to unmarshal Pod: %w", err)
		}
		return &pod, nil

	default:
		return nil, fmt.Errorf("unsupported kind: %s", api.Kind)
	}
}
