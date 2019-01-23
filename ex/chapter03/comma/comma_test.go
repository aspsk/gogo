package comma

import "testing"

func test(t *testing.T, s1, s2 string) {

	if s2 != comma_float(s1) {
		t.Errorf("%s != %s", s2, comma_float(s1))
	}

}

func TestComma(t *testing.T) {

	test(t, "123", "123")
	test(t, "123.123", "123.123")
	test(t, "-123.123", "-123.123")

	test(t, "1234", "1,234")
	test(t, "1234.123", "1,234.123")
	test(t, "-1234.123", "-1,234.123")

	test(t, "123123", "123,123")
	test(t, "123123.123", "123,123.123")
	test(t, "-123123.123", "-123,123.123")
}
