package lrotate

func lrotate(x []int, n int) {

	for n < 0 {
		n += len(x)
	}
	if n > len(x) {
		n %= len(x)
	}
	if n == 0 {
		return
	}

	y := make([]int, n)
	copy(y, x[:n])
	copy(x, x[n:])
	copy(x[len(x)-n:], y)

}
