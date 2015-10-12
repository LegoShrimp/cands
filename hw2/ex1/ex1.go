package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := new(big.Int)
	p := new(big.Int)
	p1 := new(big.Int)
	p2 := new(big.Int)
	p3 := new(big.Int)
	//Create are primes
	//SetString(string, base)
	p.SetString("906290250424699331425159556878983006056835952632974168273155634173494809116331242671104460373181196844887922027861490212777458225565043", 10)
	p1.SetString("1221045459857252130212425141209838306009872359", 10)
	p2.SetString("368939873980175183426807818564931923615802003", 10)
	p3.SetString("1005888595180246202579961831338626486331184773", 10)

	x.SetString("789897903244169481492119693927211627883741461715811732399609879256691889502496180621669114031355708742200820977788427713782862830652149", 10)
	one := new(big.Int)
	// x^(p1p2p3) = 1 mod p
	//order of x| order of p. order of p - p-1. We know the prime factors of p-1 (p11,p12,p13,2) so we just need to check those.
	//As the number is small we can just brute force it.
	//After multiple tries starting from single prime factors of phi(p) we find:
	// x^(p1p2p3) = 1 mod p
	//So the order of x is p1*p2*p3

	a := new(big.Int)
	order := new(big.Int)
	//What the next line of code does:
	//p2= p2*p3
	//order = p1*p2
	//a = x^(order)mod p
	a.Exp(x, order.Mul(p1, p2.Mul(p2, p3)), p)
	one.SetInt64(1)
	if a.String() == one.String() {

		fmt.Println("The order of x is:")
		fmt.Println(order.String())
	}

}
