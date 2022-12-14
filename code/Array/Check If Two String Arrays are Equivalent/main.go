// Author 		: Praveen
// Date   		: 14/12/2022
// Question 	: https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
// Submission 	: https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/submissions/859802257/

/*
	Question:
	--------------------------------
	Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.

	A string is represented by an array if the array elements concatenated in order forms the string.


	Example 1:
	Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
	Output: true
	Explanation:
	word1 represents string "ab" + "c" -> "abc"
	word2 represents string "a" + "bc" -> "abc"
	The strings are the same, so return true.


	Example 2:
	Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
	Output: false


*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := []string{"a", "bc"}
	word2 := []string{"ab", "c"}
	fmt.Println(arrayStringsAreEqual(word1, word2))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
