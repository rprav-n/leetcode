// Author 		: Praveen
// Date   		: 26/05/2023
// Question 	: https://leetcode.com/problems/make-the-string-great/
// Submission 	: https://leetcode.com/problems/make-the-string-great/submissions/957849154/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "kkdsFuqUfSDKK"
	fmt.Println(makeGood(s))
}

func makeGood(s string) string {

	i := 0
	for i != len(s) && i < len(s)-1 {
		c1 := rune(s[i])
		c2 := rune(s[i+1])
		if (c1 == unicode.ToLower(c2) && unicode.IsUpper(c2)) || (c2 == unicode.ToLower(c1) && unicode.IsUpper(c1)) {
			j := i + 1
			s = s[:i] + s[j+1:]
			i = 0
		} else {
			i++
		}

	}

	return s
}
