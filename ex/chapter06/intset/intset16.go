package bitset

import (
	"fmt"
	"math/bits"
)

const base16 = 16

type Bitset16 struct {
	x []uint16
}

func (bitset *Bitset16) isOn(index int) bool {

	word, bit := index/base16, uint(index%base16)
	return word < len(bitset.x) && (bitset.x[word]&(1<<bit)) != 0

}

func (bitset *Bitset16) Add(A ...int) {

	for _, index := range A {

		word, bit := index/base16, uint(index%base16)
		for len(bitset.x) <= word {
			bitset.x = append(bitset.x, 0)
		}
		bitset.x[word] |= (1 << bit)

	}

}

func (bitset *Bitset16) Remove(A ...int) {

	for _, index := range A {

		word, bit := index/base16, uint(index%base16)

		if word >= len(bitset.x) {
			continue
		}

		bitset.x[word] &^= (1 << bit)

	}

}

func (bitset *Bitset16) Clear() {
	bitset.x = []uint16{}
}

func (bitset *Bitset16) Len() int {
	var x int
	for _, i := range bitset.x {
		x += bits.OnesCount16(i)
	}
	return x
}

func (bitset *Bitset16) Copy() *Bitset16 {

	var ret *Bitset16 = &Bitset16{x: make([]uint16, len(bitset.x))}
	copy(ret.x, bitset.x)
	return ret

}

func (bitset *Bitset16) Elems() []int {

	ret := []int{}

	for i, _ := range bitset.x {
		for j := 0; j < base16; j++ {
			n := base16*i + j
			if bitset.isOn(n) {
				ret = append(ret, n)
			}
		}
	}

	return ret

}

func (bitset1 *Bitset16) UnionWith(bitset2 *Bitset16) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] |= x
		} else {
			bitset1.x = append(bitset1.x, x)
		}
	}
}

func (bitset1 *Bitset16) IntersectWith(bitset2 *Bitset16) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &= x
		}
	}
}

func (bitset1 *Bitset16) DifferenceWith(bitset2 *Bitset16) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] &^= x
		}
	}
}

func (bitset1 *Bitset16) SymmetricDifference(bitset2 *Bitset16) {
	for i, x := range bitset2.x {
		if i < len(bitset1.x) {
			bitset1.x[i] = (bitset1.x[i] &^ x) | (x &^ bitset1.x[i])
		}
	}
}

func (bitset *Bitset16) String() string {

	ret := "{ "
	for _, n := range bitset.Elems() {
		ret += fmt.Sprintf("%d ", n)
	}
	return ret + "}"
}
