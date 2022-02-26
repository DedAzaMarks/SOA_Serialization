package formats_comparison

import (
	"bytes"
	"encoding/gob"
)

type gobSerializer struct {
	b bytes.Buffer
	e *gob.Encoder
	d *gob.Decoder
}

func newGobSerializer() *gobSerializer {
	s := &gobSerializer{}
	s.e = gob.NewEncoder(&s.b)
	s.d = gob.NewDecoder(&s.b)

	if err := s.e.Encode(Test{}); err != nil {
		panic(err)
	}
	if err := s.d.Decode(&Test{}); err != nil {
		panic(err)
	}
	return s
}

func (g *gobSerializer) Marshal(t *Test) ([]byte, error) {
	g.b.Reset()
	err := g.e.Encode(t)
	return g.b.Bytes(), err
}

func (g *gobSerializer) Unmarshal(b []byte, t *Test) error {
	g.b.Reset()
	g.b.Write(b)
	err := g.d.Decode(t)
	return err
}
