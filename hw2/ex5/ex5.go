package main

import (
	"fmt"
	"github.com/legoshrimp/cands/hw2/ex2/coding"
)

func main() {
	//step 1 decrypt vigenere
	c := "jnfcztiecytrtkxbrvwgoxjvbuyducvsridemeatptnxxouhkxkhosujeezvwukwoqpekfyrvojliwojonvxijzgrjusxehs otx"
	i := coding.Encode(c)
	k := make([]int, 7)
	//Since the ordering doesn't matter because n-c-v=n-v-c
	//We can decrypt everything with the Caesar cipher first
	//We check all possible caesar ciphers. and inspect for english visually.
	for j := 0; j < 27; j++ {
		temp := shift(i, j)
		//Because we know the format of the message we can assume that k[0] will decrypt c[7]  to ' '
		//We assume k[0] will decrypt it.

		k[0] = (temp[7] - 26 + 27) % 27 //assume That k[0] will decrypt a ' '
		k[5] = (temp[0] - k[0] + 27) % 27
		k[3] = (temp[5] - k[5] + 27) % 27
		k[2] = (temp[3] - k[3] + 27) % 27
		k[4] = (temp[2] - k[2] + 27) % 27
		k[1] = (temp[4] - k[4] + 27) % 27
		k[6] = (temp[1] - k[1] + 27) % 27
		if temp[6]-k[6] == k[0] { //Verify that are assumption is correct
			fmt.Println("Good")
		}
		temp = decryptV(temp, k)

		fmt.Println(j, (temp[7]-26+27)%27, coding.Decode(temp))
	}
	// we know that pos0 and pos7 are shifted the same amount. So in all cases that pos7=' ', pos0='e'
	//phi = [5, 6, 4, 2, 1, 3, 0]
	//Encrypted key: jnfczti
	//we know that 'i' encrypted 'j', and ' '
	// DC(DV(jnfcztie))= '5642130 ' where0-6 indicate the key
	//So we know that 'e'-k6-C = ' '
	//We know that the m[7] = ' '

	//step 2 decrypt caesar cipher
	//new idea first decode the shift cipher.
	//If we do this we have
	//c' where we know that there is a space as the 8th character.

}

func decryptV(c, k []int) []int {
	m := make([]int, len(c))
	l := len(k)
	for i, v := range c {
		m[i] = (c[i] - k[i%l] + 27) % 27
		v++
	}
	return m
}
func shift(base []int, shift int) []int {
	r := make([]int, len(base))
	for i, k := range base {
		r[i] = ((k - shift) + 27) % 27
	}
	return r
}
