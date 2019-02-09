package main

import (
	"flag"
	"fmt"
	"time"
)

var imax *int = flag.Int("max", 1000, "the maximum depth")

func foo(i int, ch chan<- int) {
	if i == *imax {
		ch <- i
	} else {
		chch := make(chan int)
		go foo(i+1, chch)
		ch <- <-chch
	}
}

func main() {
	flag.Parse()
	start := time.Now()
	ch := make(chan int)
	go foo(0, ch)
	fmt.Printf("%d [in %s]\n", <-ch, time.Since(start))
}
