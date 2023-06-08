// Author 		: Praveen
// Date   		: 08/06/2023
// Question 	: https://leetcode.com/problems/first-unique-character-in-a-string/
// Submission 	: https://leetcode.com/problems/first-unique-character-in-a-string/submissions/966803907/

package main

import (
	"fmt"
	"math"
)

func main() {
	s := "aabb"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	var ans int

	m := make(map[rune]int)

	freq := make(map[rune]int)

	for i, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = i
		}

		if val, ok := freq[c]; ok {
			freq[c] = val + 1
		} else {
			freq[c] = 1
		}
	}
	ans = math.MaxInt
	isPresent := false

	for _, c := range s {
		if freq[c] == 1 {
			index := m[c]
			if index < ans {
				ans = index
			}
			isPresent = true
		}
	}
	if !isPresent {
		return -1
	}

	return ans
}
