// Author 		: Praveen
// Date   		: 28/12/2022
// Question 	: https://leetcode.com/problems/count-prefixes-of-a-given-string/
// Submission 	: https://leetcode.com/problems/count-prefixes-of-a-given-string/submissions/867055972/

/*
	Question:
	--------------------------------
	You are given a string array words and a string s, where words[i] and s comprise only of lowercase English letters.

	Return the number of strings in words that are a prefix of s.

	A prefix of a string is a substring that occurs at the beginning of the string. A substring is a contiguous sequence of characters within a string.


	Example 1:
	Input: words = ["a","b","c","ab","bc","abc"], s = "abc"
	Output: 3
	Explanation:
	The strings in words which are a prefix of s = "abc" are:
	"a", "ab", and "abc".
	Thus the number of strings in words which are a prefix of s is 3.


*/

package main

import (
	"fmt"
)

func main() {
	words := []string{"a", "b", "c", "ab", "bc", "abc"}
	s := "abc"
	fmt.Println(countPrefixes(words, s))
}

func countPrefixes(words []string, s string) int {
	var count int
	sl := len(s)
	possiblePrefixes := []string{}

	for i := 0; i < sl; i++ {
		possiblePrefixes = append(possiblePrefixes, s[:i+1])
	}

	for _, word := range words {
		for _, p := range possiblePrefixes {
			if p == word {
				count++
			}
		}
	}

	return count
}
