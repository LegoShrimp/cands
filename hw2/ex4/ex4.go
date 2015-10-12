package main

import (
	"fmt"
	"github.com/legoshrimp/cands/hw2/ex2/coding"
)

func main() {
	//Create a rune(char) slice(array) with the ciphertext
	s := []rune("apmhuqopafh mkzmahdwzlhq hjtiksjittqvo")
	//Found by visual inspection shifted by 19
	//the mighty secret word is blackballing
	//Loop that goes through all 26 possible shifts
	for i := 0; i < 27; i++ {
		//shift t i places.
		t := shift(s, i)
		fmt.Println(i, t)
	}
}

func shift(base []rune, shift int) string {
	r := make([]int, len(base))
	//Encode the string so that shifting is simple
	s := coding.Encode(string(base))
	for i, k := range s {
		//i is the index, k is s[i]
		r[i] = (k + shift) % 27
	}
	//Decode the shifted string while returning i.
	return coding.Decode(r)
}
