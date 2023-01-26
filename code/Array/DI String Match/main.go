// Author 		: Praveen
// Date   		: 26/01/2023
// Question 	: https://leetcode.com/problems/di-string-match/
// Submission 	: https://leetcode.com/problems/di-string-match/submissions/885541772/

package main

import "fmt"

func main() {
	s := "IDID"
	fmt.Println(diStringMatch(s))
}

func diStringMatch(s string) []int {
	l := len(s)

	ans := make([]int, l)

	maxI := 0
	maxD := l
	isI := false
	for i, c := range s {

		if c == rune('I') {
			ans[i] = maxI
			maxI++
			isI = true
		} else {
			ans[i] = maxD
			maxD--
			isI = false
		}
		if i == len(s)-1 {
			if isI {
				ans = append(ans, maxI)
			} else {
				ans = append(ans, maxD)
			}

		}
	}

	return ans
}
