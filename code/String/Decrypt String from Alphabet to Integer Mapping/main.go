// Author 		: Praveen
// Date   		: 27/04/2023
// Question 	: https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
// Submission 	: https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/submissions/940645606/

package main

import (
	"fmt"
)

func main() {
	s := "10#11#12"
	fmt.Println(freqAlphabets(s))
}

func freqAlphabets(s string) string {
	var ans string

	alphabets := "abcdefghijklmnopqrstuvwxyz"

	am := make(map[string]string)

	for i, c := range alphabets {
		i += 1
		st := string(c)
		if i <= 9 {
			am[fmt.Sprintf("%d", i)] = st
		} else {
			am[fmt.Sprintf("%d#", i)] = st
		}
	}

	l := len(s)
	i := l - 1
	for i >= 0 {
		str := string(s[i])
		if str == "#" {
			val := fmt.Sprintf("%s%s%s", string(s[i-2]), string(s[i-1]), "#")
			ans += am[val]
			i -= 3
		} else {
			val := fmt.Sprintf("%s", str)
			ans += am[val]
			i--
		}
	}

	ans = Reverse(ans)

	return ans
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
