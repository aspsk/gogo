package uliminate

import "testing"

func testUliminate(t *testing.T, x, y []byte) {

	x = uliminate(x)
	if string(x) != string(y) {
		t.Errorf("%v != %v", x, y)
	}

}

func TestUliminate(t *testing.T) {

	var x1, x2 string

	x1 = "Uhha Uhha Muhha Oh Oh"
	x2 = "Uhha Uhha Muhha Oh Oh"
	testUliminate(t, []byte(x1), []byte(x2))

	x1 = "Uhha 	 Uhha 	 Muhha Oh Oh 	"
	x2 = "Uhha Uhha Muhha Oh Oh "
	testUliminate(t, []byte(x1), []byte(x2))

	x1 = "			Uhha 	 Uhha 	 Muhha Oh Oh"
	x2 = "	Uhha Uhha Muhha Oh Oh"
	testUliminate(t, []byte(x1), []byte(x2))

	x1 = "Раз,  два,		три, 	четыре"
	x2 = "Раз, два,	три, четыре"
	testUliminate(t, []byte(x1), []byte(x2))

}
