package formats_comparison

import (
	"math/rand"
	"os"
	"testing"
)

type Serializer interface {
	Marshal(t *Test) ([]byte, error)
	Unmarshal(b []byte, t *Test) error
}

const dataSizeScale = 100

func benchMarshal(b *testing.B, s Serializer) {
	b.Helper()
	data := genTest(dataSizeScale)
	dir, err := os.MkdirTemp(".", "serialization")
	if err != nil {
		b.Fatalf("marshal error: %v", err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		o := data[rand.Intn(len(data))]
		bytes, err := s.Marshal(o)
		if err != nil {
			b.Fatalf("marshal error %s for %#v", err, o)
		}
		serialSize += len(bytes)
		f, err := os.CreateTemp(dir, "")
		if err != nil {
			b.Fatalf("marshal error: %v", err)
		}
		if _, err := f.Write(bytes); err != nil {
			b.Fatalf("marshal error: %v", err)
		}
		_ = f.Close()
	}
	b.StopTimer()
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func benchUnmarshal(b *testing.B, s Serializer) {
	b.Helper()
	b.StopTimer()
	dir, err := os.MkdirTemp(".", "deserialization")
	if err != nil {
		b.Fatalf("unmarshal error: %v", err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)
	data := genTest(dataSizeScale)
	var files = make([]string, len(data))
	var serialSize int
	for i, d := range data {
		o, err := s.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(o)
		file, err := os.CreateTemp(dir, "")
		if err != nil {
			b.Fatalf("unmarshal error: %v", err)
		}
		if _, err := file.Write(o); err != nil {
			b.Fatalf("unmarshal error: %v", err)
		}
		files[i] = file.Name()
		_ = file.Close()
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(files))
		bytes, err := os.ReadFile(files[n])
		if err != nil {
			b.Fatalf("unmarshal error: %v", err)
		}
		if err := s.Unmarshal(bytes, &Test{}); err != nil {
			b.Fatalf("unmarshal error %s for %#x / %q", err, bytes, bytes)
		}
	}
}

func Benchmark(b *testing.B) {
	b.Run("GOB (go-native)", func(b *testing.B) {
		g := newGobSerializer()
		b.Run("Encoding", func(b *testing.B) {
			benchMarshal(b, g)
		})
		b.Run("Decoding", func(b *testing.B) {
			benchUnmarshal(b, g)
		})
	})
	b.Run("XML", func(b *testing.B) {
		b.Run("Encoding", func(b *testing.B) {
			benchMarshal(b, xmlSerializer{})
		})
		b.Run("Decoding", func(b *testing.B) {
			benchUnmarshal(b, xmlSerializer{})
		})
	})
	b.Run("JSON", func(b *testing.B) {
		b.Run("Encoding", func(b *testing.B) {
			benchMarshal(b, jsonSerializer{})
		})
		b.Run("Decoding", func(b *testing.B) {
			benchUnmarshal(b, jsonSerializer{})
		})
	})
	b.Run("PROTOBUF", func(b *testing.B) {
		b.Run("Encoding", func(b *testing.B) {
			benchMarshal(b, protoSerializer{})
		})
		b.Run("Decoding", func(b *testing.B) {
			benchUnmarshal(b, protoSerializer{})
		})
	})
	b.Run("AVRO", func(b *testing.B) {
		b.Run("Encoding", func(b *testing.B) {
			benchMarshal(b, newAvroSerializer())
		})
		b.Run("Decoding", func(b *testing.B) {
			benchUnmarshal(b, newAvroSerializer())
		})
	})
	b.Run("YAML", func(b *testing.B) {
		b.Run("Encoding", func(b *testing.B) {
			benchMarshal(b, yamlSerializer{})
		})
		b.Run("Decoding", func(b *testing.B) {
			benchUnmarshal(b, yamlSerializer{})
		})
	})
	b.Run("MSGPACK", func(b *testing.B) {
		b.Run("Encoding", func(b *testing.B) {
			benchMarshal(b, jsonSerializer{})
		})
		b.Run("Decoding", func(b *testing.B) {
			benchUnmarshal(b, jsonSerializer{})
		})
	})
}
