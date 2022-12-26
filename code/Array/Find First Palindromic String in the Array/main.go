// Author 		: Praveen
// Date   		: 26/12/2022
// Question 	: https://leetcode.com/problems/find-first-palindromic-string-in-the-array/
// Submission 	: https://leetcode.com/problems/find-first-palindromic-string-in-the-array/submissions/865870278/

/*
	Question:
	--------------------------------
	Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

	A string is palindromic if it reads the same forward and backward.


	Example 1:
	Input: words = ["abc","car","ada","racecar","cool"]
	Output: "ada"
	Explanation: The first string that is palindromic is "ada".
	Note that "racecar" is also palindromic, but it is not the first.


*/

package main

import "fmt"

func main() {
	words := []string{"def", "ghi"}
	fmt.Println(firstPalindrome(words))
}

func firstPalindrome(words []string) string {
	var s string

loop:
	for _, word := range words {
		i := 0
		j := len(word) - 1

		for i <= j {
			if word[i] != word[j] {
				continue loop
			}
			i++
			j--
		}
		s = word
		break
	}

	return s
}
