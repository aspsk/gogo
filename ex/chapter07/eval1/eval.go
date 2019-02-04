// exercise 7.13

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/adonovan/gopl.io/ch7/eval" // should be patched to export types
)

func String(e eval.Expr) string {

	switch e := e.(type) {

	case eval.Var:
		return string(e)
	case eval.Literal:
		return fmt.Sprintf("%g", e)
	case eval.Unary:
		return string(e.Op) + String(e.X)
	case eval.Binary:
		return "(" + String(e.X) + string(e.Op) + String(e.Y) + ")"
	case eval.Call:
		ss := make([]string, len(e.Args))
		for i, arg := range e.Args {
			ss[i] = String(arg)
		}
		return fmt.Sprintf("%s(%v)", e.Fn, strings.Join(ss, ","))
	}

	log.Fatalf("unknown type %T", e)
	return ""
}

func main() {

	expr_src := "sin(1 + pow(5, 4)) / 9 * (F - 32)"

	expr, err := eval.Parse(expr_src)
	if err != nil {
		log.Fatal(err)
	}

	s := String(expr)
	fmt.Printf("expr1='%s'\n", s)

	expr2, err := eval.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("expr2='%s'\n", String(expr2))

}
