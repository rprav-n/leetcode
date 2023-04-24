// Author 		: Praveen
// Date   		: 24/04/2023
// Question 	: https://leetcode.com/problems/to-lower-case/
// Submission 	: https://leetcode.com/problems/to-lower-case/submissions/938727462/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello"
	fmt.Println(toLowerCase(s))
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}
