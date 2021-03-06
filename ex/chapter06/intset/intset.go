package bitset

import (
	"fmt"
	"math/bits"
)

const base64 = 64

type Bitset64 struct {
	x []uint64
}

func (bitset *Bitset64) isOn(index int) bool {

	word, bit := index/base64, uint(index%base64)
	return word < len(bitset.x) && (bitset.x[word]&(1<<bit)) != 0

}

func (bitset *Bitset64) Add(A ...int) {

	for _, index := range A {

		word, bit := index/base64, uint(index%base64)
		for len(bitset.x) <= word {
			bitset.x = append(bitset.x, 0)
		}
		bitset.x[word] |= (1 << bit)

	}

}

func (bitset *Bitset64) Remove(A ...int) {

	for _, index := range A {

		word, bit := index/base64, uint(index%base64)

		if word >= len(bitset.x) {
			continue
		}

		bitset.x[word] &^= (1 << bit)

	}

}

func (bitset *Bitset64) Clear() {
	bitset.x = []uint64{}
}

func (bitset *Bitset64) Len() int {
	var x int
	for _, i := range bitset.x {
		x += bits.OnesCount64(i)
	}
	return x
}

func (bitset *Bitset64) Copy() *Bitset64 {

	var ret *Bitset64 = &Bitset64{x: make([]uint64, len(bitset.x))}
	copy(ret.x, bitset.x)
	return ret

}

func (bitset *Bitset64) Elems() []int {

	ret := []int{}

	for i, _ := range bitset.x {
		for j := 0; j < base64; j++ {
			n := base64*i + j
			if bitset.isOn(n) {
				ret = append(ret, n)
			}
		}
	}

	return ret

}

func (bitset1 *Bitset64) UnionWith(bitset2 *Bitset64) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] |= x
		} else {
			bitset1.x = append(bitset1.x, x)
		}
	}
}

func (bitset1 *Bitset64) IntersectWith(bitset2 *Bitset64) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &= x
		}
	}
}

func (bitset1 *Bitset64) DifferenceWith(bitset2 *Bitset64) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &^= x
		}
	}
}

func (bitset1 *Bitset64) SymmetricDifference(bitset2 *Bitset64) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] = (bitset1.x[i] &^ x) | (x &^ bitset1.x[i])
		}
	}
}

func (bitset *Bitset64) String() string {

	ret := "{ "
	for _, n := range bitset.Elems() {
		ret += fmt.Sprintf("%d ", n)
	}
	return ret + "}"
}
