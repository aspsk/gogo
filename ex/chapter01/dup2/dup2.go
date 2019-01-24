package main

import (
	"bufio"
	"fmt"
	"os"
)

type info = struct {

	m map[string]int
	n int

}

func main() {

	counts := make(map[string]info)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "[stdin]", counts)
		files = append(files, "[stdin]")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}

	for line, i := range counts {
		if i.n > 1 {
			str := ""
			for _, file := range files {
				if i.m[file] != 0 {
					str += " " + file
				}
			}
			fmt.Printf("%d\t%s\n\t[files%s]\n", i.n, line, str)
		}
	}

}

func countLines(f *os.File, filename string, counts map[string]info) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		x := counts[input.Text()]
		if x.m == nil {
			x.m = make(map[string]int)
		}
		x.m[filename] = 1
		x.n++
		counts[input.Text()] = x
	}
}
