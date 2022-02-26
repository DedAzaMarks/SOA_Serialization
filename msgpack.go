package formats_comparison

import "github.com/shamaton/msgpack/v2"

type msgpackSerializer struct{}

func (msgpackSerializer) Marshal(t *Test) ([]byte, error) {
	return msgpack.Marshal(t)
}

func (msgpackSerializer) Unmarshal(b []byte, t *Test) error {
	return msgpack.Unmarshal(b, t)
}
