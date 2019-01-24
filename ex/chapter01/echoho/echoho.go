package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// ex 1.3: compare slow join version with strings.Join
func slowJoin(sslice []string, sep string) string {

	s, ss := "", ""
	for _, str := range sslice {
		s += ss + str
		ss = sep
	}

	return s
}

func fastJoin(sslice []string, sep string) string {
	return strings.Join(sslice, sep)
}

// ex 1.1, 1.2
func main() {

	argv0 := flag.Bool("a", false, "Do we want to print argv[0]")
	numbers := flag.Bool("n", false, "Do we want to print numbers and arguments one per line?")
	flag.Parse()

	args := flag.Args()
	if *argv0 {
		args = append([]string{os.Args[0]}, args...)
	}

	if *numbers {
		for i, arg := range args {
			fmt.Printf("%d:\t%s\n", i, arg)
		}
	} else {
		fmt.Println(strings.Join(args, " "))
	}

}
