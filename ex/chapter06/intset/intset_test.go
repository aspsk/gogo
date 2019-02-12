package bitset

import (
	"sort"
	"testing"
)

type slowBitset map[int]bool

func (sb *slowBitset) Elems() []int {
	keys := []int{}
	for key, _ := range *sb {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func compareSlices(x, y []int) bool {

	if len(x) != len(y) {
		return false
	}

	for i, xx := range x {
		if xx != y[i] {
			return false
		}
	}

	return true
}

func TestBitset(t *testing.T) {

	x := slowBitset{}
	b := Bitset64{}

	for i := 0; i < 10000; i++ {
		x[i] = true
		b.Add(i)
	}

	if !compareSlices(x.Elems(), b.Elems()) {
		t.Errorf("x.Elems=%v b.Elems=%v", x.Elems(), b.Elems())
	}

}

func BenchmarkBitset8(b *testing.B) {

	B := &Bitset8{}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			B.Add(i)
		}
	}

}

func BenchmarkBitset16(b *testing.B) {

	B := &Bitset16{}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			B.Add(i)
		}
	}

}

func BenchmarkBitset32(b *testing.B) {

	B := &Bitset32{}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			B.Add(i)
		}
	}

}

func BenchmarkBitset64(b *testing.B) {

	B := &Bitset64{}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			B.Add(i)
		}
	}

}

func BenchmarkBitsetMap(b *testing.B) {

	B := slowBitset{}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			B[i] = true
		}
	}

}
