package pod

type Metadata struct {
	Name   string            `yaml:"name"   json:"name"`
	Labels map[string]string `yaml:"labels,omitempty" json:"labels,omitempty"`
}

type Container struct {
	Name    string   `yaml:"name"      json:"name"`
	Image   string   `yaml:"image"     json:"image"`
	Command []string `yaml:"command,omitempty"   json:"command,omitempty"`
}

type Spec struct {
	Containers []*Container `yaml:"containers" json:"containers"`
}

type Pod struct {
	Kind     string   `yaml:"kind"       json:"kind"`
	Metadata Metadata `yaml:"metadata"   json:"metadata"`
	Spec     Spec     `yaml:"spec"       json:"spec"`
}
