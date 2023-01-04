// Author 		: Praveen
// Date   		: 04/01/2023
// Question 	: https://leetcode.com/problems/count-pairs-of-similar-strings/
// Submission 	: https://leetcode.com/problems/count-pairs-of-similar-strings/submissions/871251705/

/*
	Question:
	--------------------------------
	You are given a 0-indexed string array words.

	Two strings are similar if they consist of the same characters.

	For example, "abca" and "cba" are similar since both consist of characters 'a', 'b', and 'c'.
	However, "abacba" and "bcfd" are not similar since they do not consist of the same characters.
	Return the number of pairs (i, j) such that 0 <= i < j <= word.length - 1 and the two strings words[i] and words[j] are similar.


	Example 1:
	Input: words = ["aba","aabb","abcd","bac","aabc"]
	Output: 2
	Explanation: There are 2 pairs that satisfy the conditions:
	- i = 0 and j = 1 : both words[0] and words[1] only consist of characters 'a' and 'b'.
	- i = 3 and j = 4 : both words[3] and words[4] only consist of characters 'a', 'b', and 'c'.


*/

package main

import "fmt"

func main() {
	words := []string{"aabb", "ab", "ba"}

	fmt.Println(similarPairs(words))
}

func similarPairs(words []string) int {
	var count int
	l := len(words)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if Check(words[i], words[j]) {
				count++
			}
		}
	}

	return count
}

func Check(a, b string) bool {
	am := make(map[rune]struct{})
	bm := make(map[rune]struct{})

	for _, c := range a {
		am[c] = struct{}{}
	}
	for _, c := range b {
		bm[c] = struct{}{}
	}

	if len(am) != len(bm) {
		return false
	}
	c := len(am)
	cs := 0
	for key := range am {
		if _, ok := bm[key]; ok {
			cs++
		}
	}

	if cs == c {
		return true
	}

	return false
}
