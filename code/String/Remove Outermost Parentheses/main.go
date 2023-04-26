// Author 		: Praveen
// Date   		: 26/04/2023
// Question 	: https://leetcode.com/problems/remove-outermost-parentheses/
// Submission 	: https://leetcode.com/problems/remove-outermost-parentheses/submissions/940185172/

package main

import "fmt"

func main() {
	s := "(()())(())"
	fmt.Println(removeOuterParentheses(s))
}

func removeOuterParentheses(s string) string {
	var ans string

	arr := []string{}

	lp := 0
	rp := 0
	p := ""
	for _, c := range s {

		p += string(c)

		if c == rune('(') {
			lp++
		} else {
			rp++
		}

		if lp == rp && lp != 0 && rp != 0 {
			lp = 0
			rp = 0
			arr = append(arr, p)
			p = ""
		}
	}

	for _, a := range arr {
		i := 0
		j := len(a) - 1
		ans += a[i+1 : j]
	}

	return ans
}
