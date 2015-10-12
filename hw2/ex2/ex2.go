package main

import (
	"fmt"
	"github.com/legoshrimp/cands/hw2/ex2/coding"
)

func main() {
	//I think this is pretty self explanatory (s is a string, but you probably are used to this from python)
	s := "qualitatively"
	//coding.Encode returns an int slice, which is basically just an array that can be resized.
	i := coding.Encode(s)
	fmt.Println(i)
	//Initalize c as an int slice.
	c := []int{19, 7, 4, 26, 18, 4, 2, 17, 4, 19, 26, 22, 14, 17, 3, 26, 8, 18, 26, 1, 4, 11, 0, 1, 14, 17}
	//decode, which returns a string
	i2 := coding.Decode(c)
	fmt.Println(i2)

}
