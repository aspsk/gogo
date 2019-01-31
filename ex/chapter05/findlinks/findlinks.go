package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node, elements map[string]int) []string {

	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if nn := n.NextSibling; nn != nil {
		links = visit(links, nn, elements)
	}
	if nn := n.FirstChild; nn != nil {
		links = visit(links, nn, elements)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("findlinks1: %v\n", err)
	}

	elements := map[string]int{}
	links := visit(nil, doc, elements)

	fmt.Printf("---- Links ----\n")
	for _, link := range links {
		fmt.Println(link)
	}

	fmt.Printf("---- Element names ----\n")
	for key, value := range elements {
		fmt.Printf("%s: %d\n", key, value)
	}
}
