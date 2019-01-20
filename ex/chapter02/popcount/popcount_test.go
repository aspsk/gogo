package popcount

import (
	"fmt"
	"testing"
)

var testSet map[uint64]int

func init() {

	testSet := make(map[uint64]int, 0)

	testSet[0] = 0
	testSet[1] = 1
	testSet[2<<5] = 1
	testSet[3<<17] = 2
	testSet[7<<31] = 3
	testSet[0xff] = 8 * 1
	testSet[0xffff] = 8 * 2
	testSet[0xffffff] = 8 * 3
	testSet[0xffffffff] = 8 * 4
	testSet[0xffffffffff] = 8 * 5
	testSet[0xffffffffffff] = 8 * 6
	testSet[0xffffffffffffff] = 8 * 7
	testSet[0xffffffffffffffff] = 8 * 8
}

func test(foo func(x uint64) int) (uint64, error) {

	for key, _ := range testSet {
		x := foo(key)
		if x != testSet[key] {
			return key, fmt.Errorf("foo(%x) = %d (!= %d)", key, x, testSet[key])
		}
	}

	return 0, nil
}

func testt(t *testing.T, name string, foo func(x uint64) int) {

	key, err := test(foo)
	if err != nil {
		x := foo(key)
		t.Errorf("%s(%x) = %d (!= %d)", name, key, x, testSet[key])
	}

}

func TestPopCount(t *testing.T) {
	testt(t, "PopCount", PopCount)
}

func TestPopCount2(t *testing.T) {
	testt(t, "PopCount", PopCount2)
}

func TestPopCountC(t *testing.T) {
	testt(t, "PopCount", PopCountC)
}

func TestPopCountLoop(t *testing.T) {
	testt(t, "PopCountLoop", PopCountLoop)
}

func TestPopCountShift(t *testing.T) {
	testt(t, "PopCountShift", PopCountShift)
}

func TestPopCountRightmost(t *testing.T) {
	testt(t, "PopCountRightmost", PopCountRightmost)
}

func btest(N int, foo func(x uint64) int) {
	for i := 0; i < N; i++ {
		for j := 0; j < 10000000; j++ {
			foo(uint64(i))
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	btest(b.N, PopCount)
}

func BenchmarkPopCount2(b *testing.B) {
	btest(b.N, PopCount2)
}

func BenchmarkPopCountC(b *testing.B) {
	btest(b.N, PopCountC)
}

func BenchmarkPopCountC10000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountC10000000(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	btest(b.N, PopCountLoop)
}

func BenchmarkPopCountShift(b *testing.B) {
	btest(b.N, PopCountShift)
}

func BenchmarkPopCountRightmost(b *testing.B) {
	btest(b.N, PopCountRightmost)
}
