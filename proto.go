package formats_comparison

import (
	"formats-comparison/protoTest"

	"google.golang.org/protobuf/proto"
)

type protoSerializer struct {
}

func (protoSerializer) Marshal(t *Test) ([]byte, error) {
	return proto.Marshal(&protoTest.Test{
		Bool:  t.Bool,
		R:     t.R,
		I32:   t.I32,
		I64:   t.I64,
		F32:   t.F32,
		F64:   t.F64,
		S:     t.S,
		SBool: t.SBool,
		SR:    t.SR,
		SI32:  t.SI32,
		SI64:  t.SI64,
		SF32:  t.SF32,
		SF64:  t.SF64,
		SS:    t.SS,
		MS:    t.MS})
}

func (protoSerializer) Unmarshal(b []byte, t *Test) error {
	return proto.Unmarshal(b, &protoTest.Test{})
}
