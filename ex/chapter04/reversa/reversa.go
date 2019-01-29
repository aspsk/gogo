package reversa

const n = 5
func reversa(x *[n]int) {

	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-1-i] = x[len(x)-1-i], x[i]
	}

}
