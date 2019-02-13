// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
	"time"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"linear algebra":        {"calculus"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming", "compilers"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func foo(i int, course string) {
	pfx := ""
	for j := 0; j < i; j++ {
		pfx = pfx + "  "
	}
	fmt.Printf("%s%s\n", pfx, course)
	time.Sleep(100 * time.Millisecond)
}

func contains(X []string, x string) bool {
	for i := 0; i < len(X); i++ {
		if x == X[i] {
			return true
		}
	}
	return false
}

func printLoop(X []string, x string) {
	for i := 0; i < len(X); i++ {
		fmt.Printf("%s -> ", X[i])
	}
	fmt.Printf("%s\n", x)
}

func findCycles(i int, course string, stack []string, m map[string][]string) bool {

	foo(i, course)

	if contains(stack, course) {
		printLoop(stack, course)
		return true
	}

	stack = append(stack, course)

	for _, key := range m[course] {
		if key == course {
			return true
		}
		if findCycles(i+1, key, stack, m) {
			return true
		}
	}

	return false
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	for key := range prereqs {
		findCycles(0, key, []string{}, prereqs)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
