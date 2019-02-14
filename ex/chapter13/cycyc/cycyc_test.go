package cycyc

import "testing"

func TestEqual(t *testing.T) {
	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	for _, test := range []struct {
		x interface{}
		want bool
	}{
		// slice cycles
		{cycleSlice, true},
		{cyclePtr1, true},
		{cyclePtr2, true},
		{cyclePtr1, true},
		{1, false},
		{[]string{"", "2", "1"}, false},
	}{
		if Cycyc(test.x) != test.want {
			t.Errorf("Cycyc(%v) = %t", test.x, !test.want)
		}
	}
}
