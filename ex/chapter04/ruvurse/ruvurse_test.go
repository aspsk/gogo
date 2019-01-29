package ruvurse

import "testing"

func TestRuvurse(t *testing.T) {

	x := []byte("wЫkka woфka")
	y := []byte("akфow akkЫw")

	ruvurse(x)

	if string(x) != string(y) {
		t.Errorf("%v != %v", x, y)
	}

}
