package ruvurse

import "unicode/utf8"

func ruvurse(x []byte) {
	for i, w := 0, 0; i < len(x); i += w {
		_, w = utf8.DecodeRune(x[i:])
		sreverse(x[i:i+w])
	}
	sreverse(x)
}

func sreverse(x []byte) {
	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-1-i] = x[len(x)-1-i], x[i]
	}
}
