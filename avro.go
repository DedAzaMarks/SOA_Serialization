package formats_comparison

import (
	"github.com/hamba/avro"
)

type avroSerializer struct {
	schema avro.Schema
}

func newAvroSerializer() *avroSerializer {
	var res avroSerializer
	schema, err := avro.Parse(`
{
  "name": "Test",
  "type": "record",
  "namespace": "com.acme.avro",
  "fields": [
    { "name": "Bool", "type": "boolean" },
    { "name": "R", "type": "int" },
    { "name": "I32", "type": "int" },
    { "name": "I64", "type": "long" },
    { "name": "F32", "type": "float" },
    { "name": "F64", "type": "double" },
    { "name": "S", "type": "string" },
    { "name": "SBool", "type": { "type": "array", "items": "boolean" } },
    { "name": "SR", "type": { "type": "array", "items": "int" } },
    { "name": "SI32", "type": { "type": "array", "items": "int" } },
    { "name": "SI64", "type": { "type": "array", "items": "long" } },
    { "name": "SF32", "type": { "type": "array", "items": "float" } },
    { "name": "SF64", "type": { "type": "array", "items": "double" } },
    { "name": "SS", "type": { "type": "array", "items": "string" } },
    { "name": "MS", "type": { "type": "map", "values" : "string", "default": {} } }
  ]
}
`)
	if err != nil {
		panic(err)
	}
	res.schema = schema
	return &res
}

func (a avroSerializer) Marshal(t *Test) ([]byte, error) {
	return avro.Marshal(a.schema, t)
}

func (a avroSerializer) Unmarshal(b []byte, t *Test) error {
	return avro.Unmarshal(a.schema, b, t)
}
