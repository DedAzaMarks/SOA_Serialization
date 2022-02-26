package formats_comparison

import (
	yaml "gopkg.in/yaml.v2"
)

type yamlSerializer struct{}

func (yamlSerializer) Marshal(t *Test) ([]byte, error) {
	return yaml.Marshal(t)
}

func (yamlSerializer) Unmarshal(b []byte, t *Test) error {
	return yaml.Unmarshal(b, t)
}
