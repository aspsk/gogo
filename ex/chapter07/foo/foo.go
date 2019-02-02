package main

import (
	"fmt"
	"io"
	"os"
	"golang.org/x/net/html"
)

type Foo struct {
	s string
	nread int
}

func (foo *Foo) Read(p []byte) (n int, err error) {

	nleft := len(foo.s) - foo.nread
	if nleft == 0 {
		return 0, io.EOF
	}

	n = len(p)
	if n > nleft {
		n = nleft
	}

	copy(p, []byte(foo.s[foo.nread:foo.nread+n]))

	foo.nread += n
	return n, nil

}

func main() {

	foo := Foo{s: `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>title</title>
    <link rel="stylesheet" href="style.css">
    <script src="script.js"></script>
  </head>
  <body>
    <!-- page content -->
  </body>
</html>
`}

	doc, err := html.Parse(&foo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "foo: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
