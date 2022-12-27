// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/counting-words-with-a-given-prefix/
// Submission 	: https://leetcode.com/problems/counting-words-with-a-given-prefix/submissions/866274719/

/*
	Question:
	--------------------------------
	You are given an array of strings words and a string pref.

	Return the number of strings in words that contain pref as a prefix.

	A prefix of a string s is any leading contiguous substring of s.


	Example 1:
	Input: words = ["pay","attention","practice","attend"], pref = "at"
	Output: 2
	Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".


*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"leetcode", "win", "loops", "success"}
	pref := "code"
	fmt.Println(prefixCount(words, pref))
}

func prefixCount(words []string, pref string) int {
	var count int

	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			count++
		}
	}

	return count
}
