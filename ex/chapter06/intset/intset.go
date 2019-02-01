package bitset

import (
	"fmt"
	"math/bits"
)

const base = 32 << (^uint(0) >> 63)

type Bitset struct {
	x []uint
}

func (bitset *Bitset) isOn(index int) bool {

	word, bit := index/base, uint(index%base)
	return word < len(bitset.x) && (bitset.x[word]&(1<<bit)) != 0

}

func (bitset *Bitset) Add(A ...int) {

	for _, index := range A {

		word, bit := index/base, uint(index%base)
		for len(bitset.x) <= word {
			bitset.x = append(bitset.x, 0)
		}
		bitset.x[word] |= (1 << bit)

	}

}

func (bitset *Bitset) Remove(A ...int) {

	for _, index := range A {

		word, bit := index/base, uint(index%base)

		if word >= len(bitset.x) {
			continue
		}

		bitset.x[word] &^= (1 << bit)

	}

}

func (bitset *Bitset) Clear() {
	bitset.x = []uint{}
}

func (bitset *Bitset) Len() int {
	var x int
	for _, i := range bitset.x {
		x += bits.OnesCount(i)
	}
	return x
}

func (bitset *Bitset) Copy() *Bitset {

	var ret *Bitset = &Bitset{x: make([]uint, len(bitset.x))}
	copy(ret.x, bitset.x)
	return ret

}

func (bitset *Bitset) Elems() []int {

	ret := []int{}

	for i, _ := range bitset.x {
		for j := 0; j < base; j++ {
			n := base*i + j
			if bitset.isOn(n) {
				ret = append(ret, n)
			}
		}
	}

	return ret

}

func (bitset1 *Bitset) UnionWith(bitset2 *Bitset) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] |= x
		} else {
			bitset1.x = append(bitset1.x, x)
		}
	}
}

func (bitset1 *Bitset) IntersectWith(bitset2 *Bitset) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &= x
		}
	}
}

func (bitset1 *Bitset) DifferenceWith(bitset2 *Bitset) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &^= x
		}
	}
}

func (bitset1 *Bitset) SymmetricDifference(bitset2 *Bitset) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] = (bitset1.x[i] &^ x) | (x &^ bitset1.x[i])
		}
	}
}

func (bitset *Bitset) String() string {

	ret := "{ "
	for _, n := range bitset.Elems() {
		ret += fmt.Sprintf("%d ", n)
	}
	return ret + "}"
}
