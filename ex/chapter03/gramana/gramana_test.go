package gramana

import (
	"bytes"
	"math/rand"
	"testing"
	"crypto/sha256"
	"encoding/hex"
)

// returns a random string of length i
func random(n int) string {

	var b bytes.Buffer

	for ; n > 0; n -= 64 {
		r := make([]byte, 32)
		rand.Read(r)
		x := sha256.Sum256(r)
		s := hex.EncodeToString(x[:])
		if n >= 64 {
			b.WriteString(s)
		} else {
			b.WriteString(s[:n])
		}
	}

	return b.String()
}

// return a random permutation of the string @s
func permute(s string) string {

	b := make([]byte, len(s))

	perm := rand.Perm(len(s))
	for i, j := range(perm) {
		b[i] = s[j]
	}

	return string(b)
}

// returns two random anagrams
func good(n int) (string, string) {
	s1 := random(n)
	return s1, permute(s1)
}

// returns two random strings which are not anagrams (most probably)
func bad(n int) (string, string) {
	return random(n), random(n)
}

func TestGramana(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if s1, s2 := good(i+1); !gramana(s1, s2) {
			t.Errorf("'%s' and '%s' are actually anagrams", s1, s2)
		}
		if s1, s2 := bad(i+1); gramana(s1, s2) {
			t.Errorf("'%s' and '%s' are actually not anagrams", s1, s2)
		}
	}
}
