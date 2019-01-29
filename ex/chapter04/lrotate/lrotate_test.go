package lrotate

import "testing"

func cmp(x, y []int) bool {

	if len(x) != len(y) {
		return false
	}

	for i, a := range x {
		if y[i] != a {
			return false
		}
	}

	return true
}

func test(t *testing.T, x, y []int, n int) {

	lrotate(y, n)
	if !cmp(x, y) {
		t.Errorf("%v != %v", x, y)
	}

}

func TestLRotate(t *testing.T) {

	x  := []int{1, 2, 3, 4, 5}
	x1 := []int{2, 3, 4, 5, 1}
	x2 := []int{3, 4, 5, 1, 2}
	x3 := []int{4, 5, 1, 2, 3}
	x4 := []int{5, 1, 2, 3, 4}
	x5 := []int{1, 2, 3, 4, 5}

	test(t, x, x1, 4)
	test(t, x, x2, 3)
	test(t, x, x3, 2)
	test(t, x, x4, 1)
	test(t, x, x5, 0)

}
