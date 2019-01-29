package uliminate

import (
	"unicode"
	"unicode/utf8"
)

func uliminate(x []byte) []byte {

	j := 0
	for i, w := 0, 0; i < len(x); {
		var r rune
		r, w = utf8.DecodeRune(x[i:])

		// copy one rune
		copy(x[j:], x[i:i+w])
		j += w
		i += w

		// if r was the space, then skip all spaces
		for ; unicode.IsSpace(r) && (i < len(x)); {
			r, w = utf8.DecodeRune(x[i:])
			if unicode.IsSpace(r) {
				i += w
			}
		}
	}

	return x[:j]
}
