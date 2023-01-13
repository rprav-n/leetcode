// Author 		: Praveen
// Date   		: 13/01/2023
// Question 	: https://leetcode.com/problems/string-matching-in-an-array/
// Submission 	:

/*
	Question:
	--------------------------------
	Given an array of string words, return all strings in words that is a substring of another word. You can return the answer in any order.

	A substring is a contiguous sequence of characters within a string


	Example 1:
	Input: words = ["mass","as","hero","superhero"]
	Output: ["as","hero"]
	Explanation: "as" is substring of "mass" and "hero" is substring of "superhero".
	["hero","as"] is also a valid answer.


*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"blue", "green", "bu"}
	fmt.Println(stringMatching(words))
}

func stringMatching(words []string) []string {
	var ans []string

	l := len(words)

	m := make(map[string]struct{})

	for i := 0; i < l; i++ {
		f := words[i]
		for j := 0; j < l; j++ {
			if i != j {
				s := words[j]

				if strings.Contains(f, s) {
					m[s] = struct{}{}
				}
			}

		}
	}

	for k := range m {
		ans = append(ans, k)
	}

	return ans
}
