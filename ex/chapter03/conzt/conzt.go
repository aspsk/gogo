package conzt

const (

	// 63 symbols
	KB = 1000
	MB = KB * KB
	GB = KB * MB
	TB = KB * GB
	PB = KB * TB
	EB = KB * PB
	ZB = KB * EB
	YB = KB * ZB

)

const (

	// 58 symbols
	k = 1000
	kb = k
	mb = k * k
	gb = k * mb
	tb = k * gb
	pb = k * tb
	eb = k * pb
	zb = k * eb
	yb = k * zb

)

/*
 * I've also checked that we can do this with 1exx, and with a rune '<some
 * letter>', but in this case this won't be integer constants, as I wanted
 */
