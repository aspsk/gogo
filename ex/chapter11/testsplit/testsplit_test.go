package testsplit

import (
	"fmt"
	"testing"
	"strings"
)

type testset struct {
	s string
	sep string
	want int
}

func test(T *testset) error {
	words := strings.Split(T.s, T.sep)
	if got := len(words); got != T.want {
		return fmt.Errorf("Split(%q, %q) returned %d words, want %d", T.s, T.sep, got, T.want)
	}
	return nil
}

func TestSplit(t *testing.T) {
	tests := []*testset{
		&testset{ "a:b:c", ":", 3, },
		&testset{ "a?b?c", "?", 3, },
	}
	for _, T := range tests {
		if err := test(T); err != nil {
			t.Error(err)
		}
	}
}
