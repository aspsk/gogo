package nroutines

import (
	"fmt"
)

func fib(i uint64) uint64 {

	if i <= 1 {
		return i
	}
	return fib(i-2)+fib(i-1)

}

func load1(i int, pr bool) uint64 {

	ret := fib(40)
	if pr {
		fmt.Printf("[%v] %v\n", i, ret)
	}
	return ret

}

func loadN(N int, pr bool) {
	for i := 0; i < N; i++ {
		load1(i, pr)
	}
}

func load1Parallel(i int, ch chan<- struct{}, pr bool) uint64 {
	defer func() { ch <- struct{}{} }()
	return load1(i, pr)
}

func loadNParallel(N int, pr bool) {
	if pr {
		fmt.Println("starting")
		defer fmt.Println("finished")
	}
	ch := make(chan struct{})
	for i := 0; i < N; i++ {
		i := i
		go load1Parallel(i, ch, pr)
	}
	for i := 0; i < N; i++ {
		<-ch
	}
}
