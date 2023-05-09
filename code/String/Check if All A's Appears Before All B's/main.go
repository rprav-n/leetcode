// Author 		: Praveen
// Date   		: 09/05/2023
// Question 	: https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/
// Submission 	: https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/submissions/947302217/

package main

import (
	"fmt"
)

func main() {
	s := "bbb"
	fmt.Println(checkString(s))
}

func checkString(s string) bool {

	seenB := false

	for _, c := range s {
		if c == 'b' {
			seenB = true
		}
		if seenB && c == 'a' {
			return false
		}
	}

	return true
}
