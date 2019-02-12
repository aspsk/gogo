package bitset

import (
	"fmt"
	"math/bits"
)

const base8 = 8

type Bitset8 struct {
	x []uint8
}

func (bitset *Bitset8) isOn(index int) bool {

	word, bit := index/base8, uint(index%base8)
	return word < len(bitset.x) && (bitset.x[word]&(1<<bit)) != 0

}

func (bitset *Bitset8) Add(A ...int) {

	for _, index := range A {

		word, bit := index/base8, uint(index%base8)
		for len(bitset.x) <= word {
			bitset.x = append(bitset.x, 0)
		}
		bitset.x[word] |= (1 << bit)

	}

}

func (bitset *Bitset8) Remove(A ...int) {

	for _, index := range A {

		word, bit := index/base8, uint(index%base8)

		if word >= len(bitset.x) {
			continue
		}

		bitset.x[word] &^= (1 << bit)

	}

}

func (bitset *Bitset8) Clear() {
	bitset.x = []uint8{}
}

func (bitset *Bitset8) Len() int {
	var x int
	for _, i := range bitset.x {
		x += bits.OnesCount8(i)
	}
	return x
}

func (bitset *Bitset8) Copy() *Bitset8 {

	var ret *Bitset8 = &Bitset8{x: make([]uint8, len(bitset.x))}
	copy(ret.x, bitset.x)
	return ret

}

func (bitset *Bitset8) Elems() []int {

	ret := []int{}

	for i, _ := range bitset.x {
		for j := 0; j < base8; j++ {
			n := base8*i + j
			if bitset.isOn(n) {
				ret = append(ret, n)
			}
		}
	}

	return ret

}

func (bitset1 *Bitset8) UnionWith(bitset2 *Bitset8) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] |= x
		} else {
			bitset1.x = append(bitset1.x, x)
		}
	}
}

func (bitset1 *Bitset8) IntersectWith(bitset2 *Bitset8) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &= x
		}
	}
}

func (bitset1 *Bitset8) DifferenceWith(bitset2 *Bitset8) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &^= x
		}
	}
}

func (bitset1 *Bitset8) SymmetricDifference(bitset2 *Bitset8) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] = (bitset1.x[i] &^ x) | (x &^ bitset1.x[i])
		}
	}
}

func (bitset *Bitset8) String() string {

	ret := "{ "
	for _, n := range bitset.Elems() {
		ret += fmt.Sprintf("%d ", n)
	}
	return ret + "}"
}
