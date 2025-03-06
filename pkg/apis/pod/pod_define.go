package pod

type Metadata struct {
	Name   string            `yaml:"name"   json:"name"`
	Labels map[string]string `yaml:"labels,omitempty" json:"labels,omitempty"`
}

type Pod struct {
	Kind string `yaml:"kind"       json:"kind"`
}
