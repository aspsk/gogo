package deliminate

import "testing"

func cmps(x, y []string) bool {

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

func testDeliminate(t *testing.T, x, y []string) {

	x = deliminate(x)
	if !cmps(x, y) {
		t.Errorf("%v != %v", x, y)
	}

}

func TestDeliminate(t *testing.T) {

	x1 := []string{"Uhha", "Uhha", "Muhha", "Oh", "Oh"}
	x2 := []string{"Uhha", "Muhha", "Oh"}

	testDeliminate(t, x1, x2)

}
