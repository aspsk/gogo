package main

import (
	"flag"
	"fmt"
)

// I am bored to write this program, so will keep it simple

func main() {

	meters := flag.Float64("meters", 0, "Meters")
	flag.Parse()

	feet := *meters / 0.3048

	hands := feet * 3
	inches := hands * 4
	barleycorns := inches * 3
	lines := barleycorns * 4
	points := lines * 6

	yards := feet / 3
	rods := yards / 5.5
	chains := rods / 4
	furlongs := chains / 10
	statuteMiles := furlongs / 8
	leagues := statuteMiles / 3

	fmt.Printf("%.2f meters are equal to\n", *meters)

	fmt.Printf("%.2f leagues\n",      leagues)
	fmt.Printf("%.2f statuteMiles\n", statuteMiles)
	fmt.Printf("%.2f furlongs\n",     furlongs)
	fmt.Printf("%.2f chains\n",       chains)
	fmt.Printf("%.2f rods\n",         rods)
	fmt.Printf("%.2f yards\n",        yards)
	fmt.Printf("%.2f feet\n",         feet)
	fmt.Printf("%.2f hands\n",        hands)
	fmt.Printf("%.2f inches\n",       inches)
	fmt.Printf("%.2f barleycorns\n",  barleycorns)
	fmt.Printf("%.2f lines\n",        lines)
	fmt.Printf("%.2f points\n",       points)
}
