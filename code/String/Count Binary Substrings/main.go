// Author 		: Praveen
// Date   		: 19/05/2023
// Question 	: https://leetcode.com/problems/count-binary-substrings/
// Submission 	: https://leetcode.com/problems/count-binary-substrings/submissions/953501644/

package main

import "fmt"

func main() {
	s := "00011100"
	fmt.Println(countBinarySubstrings(s))
}

func countBinarySubstrings(s string) int {
	count := 0
	prevCount := 0
	currCount := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			currCount++
		} else {
			count += min(prevCount, currCount)
			prevCount = currCount
			currCount = 1
		}
	}

	// Add the count of the last group
	count += min(prevCount, currCount)

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
