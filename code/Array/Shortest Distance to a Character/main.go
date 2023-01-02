// Author 		: Praveen
// Date   		: 02/01/2023
// Question 	: https://leetcode.com/problems/shortest-distance-to-a-character/
// Submission 	: https://leetcode.com/problems/shortest-distance-to-a-character/submissions/869394689/

/*
	Question:
	--------------------------------
	Given a string s and a character c that occurs in s, return an array of integers answer where answer.length == s.length and answer[i] is the distance from index i to the closest occurrence of character c in s.

	The distance between two indices i and j is abs(i - j), where abs is the absolute value function.


	Example 1:
	Input: s = "loveleetcode", c = "e"
	Output: [3,2,1,0,1,0,0,1,2,2,1,0]
	Explanation: The character 'e' appears at indices 3, 5, 6, and 11 (0-indexed).
	The closest occurrence of 'e' for index 0 is at index 3, so the distance is abs(0 - 3) = 3.
	The closest occurrence of 'e' for index 1 is at index 3, so the distance is abs(1 - 3) = 2.
	For index 4, there is a tie between the 'e' at index 3 and the 'e' at index 5, but the distance is still the same: abs(4 - 3) == abs(4 - 5) = 1.
	The closest occurrence of 'e' for index 8 is at index 6, so the distance is abs(8 - 6) = 2.


*/

package main

import (
	"fmt"
)

func main() {
	s := "aaab"
	c := byte('b')
	fmt.Println(shortestToChar(s, c))
}

func shortestToChar(s string, c byte) []int {
	var ans []int

	var indices []int

	for i, sc := range s {
		if byte(sc) == c {
			indices = append(indices, i)
		}
	}

	for i, sc := range s {
		var min int
		if byte(sc) == c {
			ans = append(ans, 0)
			continue
		}
		for j := range indices {
			ab := abs(i - indices[j])
			if ab < min || min == 0 {
				min = ab
			}
		}
		ans = append(ans, min)
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
