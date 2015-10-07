package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := make(map[string]bool)
	var x *big.Int
	x = new(big.Int)
	p := new(big.Int)
	p1 := new(big.Int)
	p2 := new(big.Int)
	p3 := new(big.Int)
	p.SetString("906290250424699331425159556878983006056835952632974168273155634173494809116331242671104460373181196844887922027861490212777458225565043", 10)
	p1.SetString("1221045459857252130212425141209838306009872359", 10)
	p2.SetString("368939873980175183426807818564931923615802003", 10)
	p3.SetString("1005888595180246202579961831338626486331184773", 10)

	x.SetString("789897903244169481492119693927211627883741461715811732399609879256691889502496180621669114031355708742200820977788427713782862830652149", 10)
	m[x.String()] = true
	n := new(big.Int)
	n.Set(x)
	one := new(big.Int)
	// x^(p1p2p3) = 1 mod p
	fmt.Println(p1.Exp(n, p1.Add(p1, big.NewInt(-1)), p))
	fmt.Println()
	fmt.Println(p2.Exp(n, p2.Add(p2, big.NewInt(-1)), p))
	fmt.Println()
	fmt.Println(p3.Exp(n, p3.Add(p3, big.NewInt(-1)), p))
	fmt.Println()
	fmt.Println(n.Exp(n, big.NewInt(2), p))
	fmt.Println()

	//k := 0
	one.SetInt64(1)
	if p1.String() == p.String() {

		fmt.Println("Good! That is whats supposed to happen!")
	}
	fmt.Println(n)

	fmt.Println(x)

}
