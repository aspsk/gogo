package bitset

import (
	"fmt"
	"math/bits"
)

const base32 = 32

type Bitset32 struct {
	x []uint32
}

func (bitset *Bitset32) isOn(index int) bool {

	word, bit := index/base32, uint(index%base32)
	return word < len(bitset.x) && (bitset.x[word]&(1<<bit)) != 0

}

func (bitset *Bitset32) Add(A ...int) {

	for _, index := range A {

		word, bit := index/base32, uint(index%base32)
		for len(bitset.x) <= word {
			bitset.x = append(bitset.x, 0)
		}
		bitset.x[word] |= (1 << bit)

	}

}

func (bitset *Bitset32) Remove(A ...int) {

	for _, index := range A {

		word, bit := index/base32, uint(index%base32)

		if word >= len(bitset.x) {
			continue
		}

		bitset.x[word] &^= (1 << bit)

	}

}

func (bitset *Bitset32) Clear() {
	bitset.x = []uint32{}
}

func (bitset *Bitset32) Len() int {
	var x int
	for _, i := range bitset.x {
		x += bits.OnesCount32(i)
	}
	return x
}

func (bitset *Bitset32) Copy() *Bitset32 {

	var ret *Bitset32 = &Bitset32{x: make([]uint32, len(bitset.x))}
	copy(ret.x, bitset.x)
	return ret

}

func (bitset *Bitset32) Elems() []int {

	ret := []int{}

	for i, _ := range bitset.x {
		for j := 0; j < base32; j++ {
			n := base32*i + j
			if bitset.isOn(n) {
				ret = append(ret, n)
			}
		}
	}

	return ret

}

func (bitset1 *Bitset32) UnionWith(bitset2 *Bitset32) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] |= x
		} else {
			bitset1.x = append(bitset1.x, x)
		}
	}
}

func (bitset1 *Bitset32) IntersectWith(bitset2 *Bitset32) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &= x
		}
	}
}

func (bitset1 *Bitset32) DifferenceWith(bitset2 *Bitset32) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &^= x
		}
	}
}

func (bitset1 *Bitset32) SymmetricDifference(bitset2 *Bitset32) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] = (bitset1.x[i] &^ x) | (x &^ bitset1.x[i])
		}
	}
}

func (bitset *Bitset32) String() string {

	ret := "{ "
	for _, n := range bitset.Elems() {
		ret += fmt.Sprintf("%d ", n)
	}
	return ret + "}"
}
