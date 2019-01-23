package comma

import (
	"bytes"
	"strings"
)

// the original comma
func comma(s string) string {

	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]

}

// 3.10: non-recursive version of comma using bytes.Buffer
func comma_nr(s string) string {

	n := len(s)
	if n <= 3 {
		return s
	}

	var b bytes.Buffer
	b.Grow(n * 4 / 3)

	// "12345678" n=8 = 6 + 2
	// "12,345,678"
	m := n % 3
	if m != 0 {
		b.WriteString(s[:m])
		b.WriteRune(',')
		n -= m
	}
	for i := 0; 3*i < n; i++ {
		b.WriteString(s[m+3*i : m+3*i+3])
		if 3*i+3 < n {
			b.WriteRune(',')
		}
	}

	return b.String()
}

func parse(s string) (string, string, string) {

	if s == "" {
		return "", "", ""
	}

	sign := 0
	if s[0] == '-' || s[0] == '+' {
		sign = 1
	}

	point := strings.Index(s, ".")
	if point == -1 {
		point = len(s)
	}

	return s[:sign], s[sign:point], s[point:]
}

// 3.11: support floating points and signs
func comma_float(s string) string {

	sign, base, tail := parse(s)
	return sign + comma_nr(base) + tail

}
