package reversa

import "testing"

func TestReversa(t *testing.T) {

	x := [n]int{1, 2, 3, 4, 5}
	y := [n]int{5, 4, 3, 2, 1}

	reversa(&x)

	if x != y {
		t.Errorf("%v != %v", x, y)
	}

}
