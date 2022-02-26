package formats_comparison

import "encoding/json"

type jsonSerializer struct{}

func (jsonSerializer) Marshal(t *Test) ([]byte, error) {
	return json.Marshal(t)
}

func (jsonSerializer) Unmarshal(b []byte, t *Test) error {
	return json.Unmarshal(b, t)
}
