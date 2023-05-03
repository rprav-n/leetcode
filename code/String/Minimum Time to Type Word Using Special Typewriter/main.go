// Author 		: Praveen
// Date   		: 03/05/2023
// Question 	: https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/
// Submission 	: https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/submissions/943922663/

package main

import (
	"fmt"
)

func main() {
	word := "pdy"
	fmt.Println(minTimeToType(word))
}

// Solution from : https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/solutions/2069859/golang-solution-0ms-100/?languageTags=golang
func minTimeToType(word string) int {
	ans := len(word) + calcMove(byte('a'), word[0])
	for i := 1; i < len(word); i++ {
		ans += calcMove(word[i-1], word[i])
	}
	return ans
}

func calcMove(a, b byte) int {
	if a > b {
		a, b = b, a
	}
	i := int(b - a)
	j := int(a + byte(26) - b)
	return min(i, j)
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
