// ex 7.15

package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"github.com/adonovan/gopl.io/ch7/eval"
)

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format + "\n", args...)
	os.Exit(1)
}

func main() {

	if len(os.Args) < 2 {
		die("usage: %s <expr> [ <arg1> ... ]", os.Args[0])
	}

	str := os.Args[1]
	expr, err := eval.Parse(str)
	if err != nil {
		die("can't parse '': %v", str, err)
	}

	env := make(map[eval.Var]float64)
	for _, arg := range os.Args[2:] {
		ss := strings.Split(arg, "=")
		if len(ss) != 2 {
			die("bad expression: '%s', expected <var>=<value>", arg)
		}
		name := eval.Var(ss[0])
		value, err := strconv.ParseFloat(ss[1], 64)
		if err != nil {
			die("bad float value: '%s'", ss[1])
		}
		env[eval.Var(name)] = value
	}

	fmt.Println(env)

	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		die("bad expression: %s: check failed: %v", str, err)
	}

	fmt.Println(vars)

	for k, _ := range vars {
		if _, ok := env[k]; !ok {
			die("variable '%s' undefined", string(k))
		}
	}

	fmt.Printf("result=%g\n", expr.Eval(env))
}
