package stringtree

import (
	"fmt"
	"bytes"
)

type tree struct {
	value       int
	left, right *tree
}

func printTree(t *tree, buf *bytes.Buffer, pfx string) {

	if t != nil {
		figul := ""
		if pfx != "" {
			figul = "└──"
		}
		pp := pfx
		if pp == "" {
			pp = "   "
		}
		fmt.Fprintf(buf, "%s%s%d\n", pp, figul, t.value)
		printTree(t.left, buf, pfx + "   ")
		printTree(t.right, buf, pfx + "   ")
	}

}

func (t *tree) String() string {
	var buf bytes.Buffer
	printTree(t, &buf, "")
	return buf.String()
}
