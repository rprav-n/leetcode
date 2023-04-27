// Author 		: Praveen
// Date   		: 27/04/2023
// Question 	: https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
// Submission 	: https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/submissions/940501266/

package main

import (
	"fmt"
	"strings"
)

func main() {
	patters := []string{"a", "abc", "bc", "d"}
	word := "abc"

	fmt.Println(numOfStrings(patters, word))
}

func numOfStrings(patterns []string, word string) int {
	var count int

	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			count++
		}
	}

	return count
}
