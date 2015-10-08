package main

import (
	"fmt"
	"github.com/legoshrimp/cands/hw2/ex2/coding"
)

func main() {
	s := "abcdedfhijklmnopqrstuvwxyz "
	i := coding.Encode(s)
	fmt.Println(i)
	i2 := coding.Decode(i)
	fmt.Println(i2)

}
