package main

import (
	"fmt"
	"github.com/legoshrimp/cands/hw2/ex2/coding"
)

//our prime P
const P = 1051

func main() {
	//initialize x with our seed values
	x := []int{341, 817, 38, 471, 50, 277, 43, 491, 796, 812}
	//Initialize c with our ciphertext
	c := []int{346, 825, 55, 489, 69, 303, 60, 511, 807, 816, 299, 556, 142, 165, 459, 605, 76, 491, 997, 182, 144, 548, 991, 788, 222, 584, 410, 439, 403, 755, 986, 91, 990, 719, 833, 715, 611, 890, 78, 436, 63, 220, 32, 32, 49, 865, 909, 420, 60}
	for i, k := range c {
		//Decode our ciphertext
		c[i] = (P + k - x[0]) % P
		//Increment our shift register
		x = shiftReg(x)
	}
	fmt.Println(coding.Decode(c))

}

//Function that performs the shift operations.
func shiftReg(x []int) []int {
	//Calculate y according to r3
	y := (415*x[9] + 542*x[8] + 1049*x[7] + 188*x[6] + 663*x[5] + 587*x[4] + 354*x[3] + 58*x[2] + 337*x[1] + 450*x[0]) % P
	x[0] = x[1]
	x[1] = x[2]
	x[2] = x[3]
	x[3] = x[4]
	x[4] = x[5]
	x[5] = x[6]
	x[6] = x[7]
	x[7] = x[8]
	x[8] = x[9]
	x[9] = y
	return x

}
