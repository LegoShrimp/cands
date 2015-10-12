package coding

import ()

func Encode(s string) []int {
	//Creates an int slice(resizable vector) with the same size as string s
	code := make([]int, len(s))
	for i, k := range s {
		//Space is a special case, if it is space it is 26
		if k == ' ' {
			code[i] = 26
		} else {
			//Otherwise we can just subtract the character 'a' to bring it into the correct range
			//as a-z are encoded sequentialy
			code[i] = int(k - 'a')
		}
	}
	return code
}
func Decode(i []int) string {
	s := make([]rune, len(i))
	for i, k := range i {
		//Same process as encoding, just if 26, and add 'a'
		if k == 26 {
			s[i] = (' ')
		} else {
			s[i] = rune(k + 'a')
		}
	}
	return string(s)
}
