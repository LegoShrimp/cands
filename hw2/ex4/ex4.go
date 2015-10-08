package main

import (
	"fmt"
	"github.com/legoshrimp/cands/hw2/ex2/coding"
)

func main() {
	s := []rune("apmhuqopafh mkzmahdwzlhq hjtiksjittqvo")
	//Found by visual inspection shifted by 19
	//the mighty secret word is blackballing
	for i := 19; i < 20; i++ {
		t := shift(s, i)
		fmt.Println(i, string(t))
	}
}

func shift(base []rune, shift int) string {
	r := make([]int, len(base))
	s := coding.Encode(string(base))
	for i, k := range s {
		r[i] = (k + shift) % 27
	}
	return coding.Decode(r)
}
