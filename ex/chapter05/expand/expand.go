package main

import (
	"fmt"
)

func ischarnum(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func expand(s string, f func(string)string) string {
	r := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '$' {
			j := i + 1
			for ; j < len(s) && ischarnum(rune(s[j])); j++ { }
			r = append(r, []byte(f(s[i+1:j]))...)
			i = j - 1
		} else {
			r = append(r, s[i])
		}
	}
	return string(r)
}

func main() {

	s := "$1$2a|2$1$$"
	fmt.Printf("%s\n%s\n", s, expand(s, func(s string) string { return s + s }))

}
