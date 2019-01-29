package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
)

type counter = struct {
	foo func(rune) bool
	n   int
}

var counters map[string]counter

func init() {

	counters = make(map[string]counter)

	counters["rune"]	= counter{foo: func(rune) bool { return true }}
	counters["control"]	= counter{foo: unicode.IsControl}
	counters["digit"]	= counter{foo: unicode.IsDigit}
	counters["graphic"]	= counter{foo: unicode.IsGraphic}
	counters["letter"]	= counter{foo: unicode.IsLetter}
	counters["lower"]	= counter{foo: unicode.IsLower}
	counters["upper"]	= counter{foo: unicode.IsUpper}
	counters["number"]	= counter{foo: unicode.IsNumber}
	counters["print"]	= counter{foo: unicode.IsPrint}
	counters["punct"]	= counter{foo: unicode.IsPunct}
	counters["space"]	= counter{foo: unicode.IsSpace}

}

func count(r rune) {

	for key, counter := range counters {
		if counter.foo(r) {
			counter.n++
			counters[key] = counter
		}
	}

}

func main() {
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		count(r)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	keys := []string{}
	for key, _ := range counters {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		n := counters[key].n
		if n > 0 {
			fmt.Printf("\n%d %ss", n, key)
		}
	}
}
