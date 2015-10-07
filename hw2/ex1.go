package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := make(map[*big.Int]bool)
	var x *big.Int
	x = new(big.Int)
	p := new(big.Int)
	p.SetString("906290250424699331425159556878983006056835952632974168273155634173494809116331242671104460373181196844887922027861490212777458225565043", 10)
	m[(x)] = true
	n := new(big.Int)
	n.Set(x)
	one := new(big.Int)
	k := 0
	one.SetInt64(1)
	for n != one || k < 1000 {
		n = n.Mul(n, x)
		n = n.Mod(n, p)
		m[(n)] = true
		fmt.Println(n)
		fmt.Println(len(m))
		k = k + 1
	}
	x.SetString("789897903244169481492119693927211627883741461715811732399609879256691889502496180621669114031355708742200820977788427713782862830652149", 10)

	fmt.Println(x)

}
