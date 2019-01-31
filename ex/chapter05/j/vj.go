package vj

import "strings"

func VJoin(sep string, a ...string) string {
	return strings.Join(a, sep)
}
