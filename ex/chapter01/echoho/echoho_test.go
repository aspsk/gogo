package main

import (
	"testing"
	"math/rand"
	"encoding/hex"
	"crypto/sha256"
)

func TestJoin(t *testing.T) {

	sep := "wokka"
	slice := []string{"once", "upon", "a", "time"}

	if s1, s2 := slowJoin(slice, sep), fastJoin(slice, sep); s1 != s2 {
		t.Errorf("%s != %s", s1, s2)
	}

}

type testStruct = struct {
	slice []string
	sep string
}

const testListSize = 1000

// test lists of testListSize elements each; the difference is in length of slices
var testList100   [testListSize]testStruct
var testList1000  [testListSize]testStruct

func randomString() string {

	var r [32]byte
	rand.Read(r[:])
	x := sha256.Sum256(r[:])
	return hex.EncodeToString(x[:])

}

func randomStringSlice(n int) []string {
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		slice[i] = randomString()
	}
	return slice
}

func init() {

	for i := 0; i < testListSize; i++ {
		testList100[i].slice = randomStringSlice(100)
		testList100[i].sep = randomString()[:2]

		testList1000[i].slice = randomStringSlice(1000)
		testList1000[i].sep = randomString()[:2]
	}

}

func doTest(foo func ([]string, string) string, testList *[testListSize]testStruct, N int) {

	for i := 0; i < N; i++ {
		rand.Seed(int64(i))
		test := testList[i % testListSize]
		foo(test.slice, test.sep)
	}

}

func BenchmarkSlowJoin100(b *testing.B) { doTest(slowJoin, &testList100, b.N) }
func BenchmarkFastJoin100(b *testing.B) { doTest(fastJoin, &testList100, b.N) }
func BenchmarkSlowJoin1000(b *testing.B) { doTest(slowJoin, &testList1000, b.N) }
func BenchmarkFastJoin1000(b *testing.B) { doTest(fastJoin, &testList1000, b.N) }
