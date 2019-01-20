package popcount

var pc [256]byte
var pc2 [65536]byte

func init() {

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	for i := range pc2 {
		pc2[i] = pc2[i/2] + byte(i&1)
	}

}

func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])

}

func PopCount2(x uint64) int {

	return int(pc2[uint16(x>>(0*8))] + pc2[uint16(x>>(2*8))] + pc2[uint16(x>>(4*8))] + pc2[uint16(x>>(6*8))])

}

func PopCountLoop(x uint64) (ret int) {

	for i := uint(0); i < 8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}

	return ret

}

func PopCountShift(x uint64) (ret int) {

	for i := uint(0); i < 64; i++ {
		ret += int(x & 1)
		x >>= 1
	}

	return ret

}

func PopCountRightmost(x uint64) (ret int) {

	for ; x != 0; x = x & (x - 1) {
		ret++
	}

	return ret
}
