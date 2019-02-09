package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCC(t *testing.T) {

	b := bytes.Buffer{}

	doit(strings.NewReader("123"), &b)

	if b.String() != `
3 digits
3 graphics
3 numbers
3 prints
3 runes` {
	t.Error("oops")
}

}
