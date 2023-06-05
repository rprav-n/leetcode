// Author 		: Praveen
// Date   		: 05/06/2023
// Question 	: https://leetcode.com/problems/reverse-only-letters/
// Submission 	: https://leetcode.com/problems/reverse-only-letters/submissions/964588347/

package main

import (
	"fmt"
)

func main() {
	s := "7_28]"
	fmt.Println(reverseOnlyLetters(s))
}

func reverseOnlyLetters(s string) string {
	var ans string
	l := len(s)
	b := []byte(s)
	start := 0
	end := l - 1

	for start < end {
		first := s[start]
		second := s[end]

		if first < 65 || (first > 90 && first < 97) || first > 122 {
			start++
			continue
		}
		if second < 65 || (second > 90 && second < 97) || second > 122 {
			end--
			continue
		}

		b[start], b[end] = b[end], b[start]
		start++
		end--

	}
	ans = string(b)
	return ans
}
