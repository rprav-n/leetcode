// Author 		: Praveen
// Date   		: 24/04/2023
// Question 	: https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
// Submission 	: https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/submissions/938735346/

package main

import (
	"fmt"
)

func main() {
	s := "(1)+((2))+(((3)))"
	fmt.Println(maxDepth(s))
}

func maxDepth(s string) int {
	var depth int
	counter := 0
	for _, c := range s {
		if c == '(' {
			counter++
		} else if c == ')' {
			counter--
		}
		depth = max(depth, counter)
	}
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
