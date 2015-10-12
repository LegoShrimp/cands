package main

import (
	"fmt"
	"github.com/legoshrimp/cands/hw2/ex2/coding"
)

func main() {
	//Initalize our ciphertext
	c := "jnfcztiecytrtkxbrvwgoxjvbuyducvsridemeatptnxxouhkxkhosujeezvwukwoqpekfyrvojliwojonvxijzgrjusxehs otx"
	//Encode our ciphertext
	i := coding.Encode(c)
	//make the key slice
	k := make([]int, 7)
	//Since the ordering doesn't matter because n-c-v=n-v-c
	//We can decrypt everything with the Caesar cipher first
	//We check all possible caesar ciphers. then visually inspect for english visually.
	//When I wrote this I wasn't sure it was safe to assume caeser was always a shift by 3 as most of what I saw when I googled to double check was using it more generally as simple shift ciphers.
	for j := 0; j < 27; j++ {
		temp := shift(i, j)
		//Because we know some of the format of the message we can assume that k[0] will decrypt c[7]  to ' '
		//So we can just hardcode the value.
		//After we know one decrpytion, it is simple to decrypt each key based on our knowledge of the key permutation.
		//phi = [5, 6, 4, 2, 1, 3, 0]

		k[0] = (temp[7] - 26 + 27) % 27 //assume That k[0] will decrypt a ' '
		k[5] = (temp[0] - k[0] + 27) % 27
		k[3] = (temp[5] - k[5] + 27) % 27
		k[2] = (temp[3] - k[3] + 27) % 27
		k[4] = (temp[2] - k[2] + 27) % 27
		k[1] = (temp[4] - k[4] + 27) % 27
		k[6] = (temp[1] - k[1] + 27) % 27
		if temp[6]-k[6] == k[0] { //Verify that our assumption is correct

			fmt.Println("Good")
		}
		temp = decryptV(temp, k)

		fmt.Println(j, (temp[7]-26+27)%27, coding.Decode(temp))
	}

}

//A function that decrypts a Viginere cipher with key k, can work for any key length.
func decryptV(c, k []int) []int {
	m := make([]int, len(c))
	l := len(k)
	for i, v := range c {
		m[i] = (c[i] - k[i%l] + 27) % 27
		v++
	}
	return m
}

//Modified shift from ex4. Now returns an int slice.
func shift(base []int, shift int) []int {
	r := make([]int, len(base))
	for i, k := range base {
		r[i] = ((k - shift) + 27) % 27
	}
	return r
}
