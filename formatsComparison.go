package formats_comparison

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Test struct {
	Bool bool    // 1
	R    rune    // 4
	I32  int32   // 4
	I64  int64   // 8
	F32  float32 // 4
	F64  float64 // 8
	S    string

	SBool []bool
	SR    []rune
	SI32  []int32
	SI64  []int64
	SF32  []float32
	SF64  []float64
	SS    []string

	MS StringMap
}

func (t *Test) SizeOf() int {
	res := 29
	res += len(t.S)

	res += len(t.SBool) * 1
	res += len(t.SR) * 4
	res += len(t.SI32) * 4
	res += len(t.SI64) * 8
	res += len(t.SF32) * 4
	res += len(t.SF64) * 8
	for _, s := range t.SS {
		res += len(s)
	}

	for key, value := range t.MS {
		res += len(key)
		res += len(value)
	}
	return res
}

func genTest(n int) []*Test {
	res := make([]*Test, 0, 1000)
	for i := 0; i < 1000; i++ {
		t := Test{
			Bool: rand.Intn(2) == 1,
			R:    rune(rand.Uint32()),
			I32:  rand.Int31(),
			I64:  rand.Int63(),
			F32:  rand.Float32(),
			F64:  rand.Float64(),
			S:    RandStringRunes(rand.Intn(n)),

			SBool: make([]bool, rand.Intn(n)),
			SR:    make([]rune, rand.Intn(n)),
			SI32:  make([]int32, rand.Intn(n)),
			SI64:  make([]int64, rand.Intn(n)),
			SF32:  make([]float32, rand.Intn(n)),
			SF64:  make([]float64, rand.Intn(n)),
			SS:    make([]string, rand.Intn(n)),

			MS: make(map[string]string),
		}

		for i := range t.SBool {
			t.SBool[i] = rand.Intn(2) == 1
		}
		for i := range t.SR {
			t.SR[i] = rune(rand.Uint32())
		}
		for i := range t.SI32 {
			t.SI32[i] = rand.Int31()
		}
		for i := range t.SI64 {
			t.SI64[i] = rand.Int63()
		}
		for i := range t.SF32 {
			t.SF32[i] = rand.Float32()
		}
		for i := range t.SF64 {
			t.SF64[i] = rand.Float64()
		}
		for i := range t.SS {
			t.SS[i] = RandStringRunes(rand.Intn(n))
		}

		for i := 0; i < rand.Intn(n); i++ {
			t.MS[RandStringRunes(1+rand.Intn(n))] = RandStringRunes(rand.Intn(n))
		}
		res = append(res, &t)
	}
	return res
}
