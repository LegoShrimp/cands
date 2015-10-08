package coding

import ()

func Encode(s string) []int {
	code := make([]int, len(s))
	for i, k := range s {
		if k == ' ' {
			code[i] = 26
		} else {
			code[i] = int(k - 'a')
		}
	}
	return code
}
func Decode(i []int) string {
	s := make([]rune, len(i))
	for i, k := range i {
		if k == 26 {
			s[i] = (' ')
		} else {
			s[i] = rune(k + 'a')
		}
	}
	return string(s)
}
