// Author 		: Praveen
// Date   		: 17/12/2022
// Question 	: https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Submission 	: https://leetcode.com/problems/count-the-number-of-consistent-strings/submissions/861204834/

/*
	Question:
	--------------------------------
	You are given a string allowed consisting of distinct characters and an array of strings words. A string is consistent if all characters in the string appear in the string allowed.

	Return the number of consistent strings in the array words.


	Example 1:
	Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
	Output: 2
	Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.

	Example 2:
	Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
	Output: 7
	Explanation: All strings are consistent.



*/

package main

import "fmt"

func main() {
	allowed := "abc"
	words := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string) int {
	var count int

	allowedMap := make(map[rune]struct{})

	for _, c := range allowed {
		allowedMap[c] = struct{}{}
	}

outerLoop:
	for _, word := range words {
		for _, c := range word {
			if _, ok := allowedMap[c]; !ok {
				continue outerLoop
			}
		}
		count++
	}

	return count
}
