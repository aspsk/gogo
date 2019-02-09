package main

import (
	"fmt"
	"time"
)

func ping(ch chan int, i *int) {
	for {
		ch <- *i
		*i++
		<-ch
	}
}

func pong(ch chan int) {
	for {
		i := <-ch
		ch <- i
	}
}

func main() {

	ch, i := make(chan int), 0

	go ping(ch, &i)
	go pong(ch)

	time.Sleep(1 * time.Second)
	fmt.Printf("i=%d\n", i)
}
