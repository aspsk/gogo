package main

import "fmt"

func min(x ...int) int {
	if len(x) == 0 {
		panic("hug this")
	}
	return min1(x[0], x[1:]...)
}

func min1(x0 int, x ...int) int {
	ret := x0
	for _, y := range x {
		if ret > y {
			ret = y
		}
	}
	return ret
}

func main() {

	fmt.Println(min(0, 1, 2))
	fmt.Println(min1(0, 1, 2))
	fmt.Println(min1(0))

}
