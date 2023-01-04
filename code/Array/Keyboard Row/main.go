// Author 		: Praveen
// Date   		: 04/01/2023
// Question 	: https://leetcode.com/problems/keyboard-row/
// Submission 	: https://leetcode.com/problems/keyboard-row/submissions/871260939/

/*
	Question:
	--------------------------------
	Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.

	In the American keyboard:

	the first row consists of the characters "qwertyuiop",
	the second row consists of the characters "asdfghjkl", and
	the third row consists of the characters "zxcvbnm".


	Example 1:
	Input: words = ["Hello","Alaska","Dad","Peace"]
	Output: ["Alaska","Dad"]


*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}

func findWords(words []string) []string {
	var ans []string

	r1 := "qwertyuiop"
	r2 := "asdfghjkl"
	r3 := "zxcvbnm"

	for _, word := range words {
		if Check(word, r1) || Check(word, r2) || Check(word, r3) {
			ans = append(ans, word)
		}
	}

	return ans
}

func Check(word string, r string) bool {
	word = strings.ToLower(word)
	for _, c := range word {
		if !strings.ContainsRune(r, c) {
			return false
		}
	}

	return true
}
