// ex 7.15

package main

import (
	"log"
	"net/http"
	"fmt"
	"strconv"
	"os"
	"github.com/adonovan/gopl.io/ch7/eval"
)

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format + "\n", args...)
	os.Exit(1)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "\n")

	expr_str, ok := r.Form["expr"]
	if !ok {
		fmt.Fprintf(w, "Expression is empty\n")
		return
	}

	expr, err := eval.Parse(expr_str[0])
	if err != nil {
		fmt.Fprintf(w, "Expression is wrong: %q: %v\n", expr_str[0], err)
		return
	}

	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		fmt.Fprintf(w, "Expression is wrong: %q: check failed: %v\n", expr_str[0], err)
		return
	}

	env := make(map[eval.Var]float64)
	for k, v := range r.Form {
		if k == "expr" {
			continue
		}

		name := eval.Var(k)
		value, err := strconv.ParseFloat(v[0], 64)
		if err != nil {
			fmt.Fprintf(w, "bad float value: '%s'\n", v[0])
			return
		}
		env[eval.Var(name)] = value
	}

	for k, _ := range vars {
		OK := true
		if _, ok := env[k]; !ok {
			fmt.Fprintf(w, "variable '%s' undefined", string(k))
			OK = false
		}
		if !OK {
			return
		}
	}

	fmt.Fprintf(w, "RESULT=%g\n", expr.Eval(env))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
