// Author 		: Praveen
// Date   		: 25/04/2023
// Question 	: https://leetcode.com/problems/reverse-words-in-a-string-iii/
// Submission 	: https://leetcode.com/problems/reverse-words-in-a-string-iii/submissions/939637731/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	var ans string
	rs := ""
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		str := string(s[i])
		rs += str
	}

	arr := strings.Split(rs, " ")

	la := len(arr)

	for i := la - 1; i >= 0; i-- {
		ans += arr[i]
		if i != 0 {
			ans += " "
		}
	}

	return ans
}
