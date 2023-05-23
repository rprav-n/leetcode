// Author 		: Praveen
// Date   		: 23/05/2023
// Question 	: https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent/
// Submission 	: https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent/submissions/955908266/

package main

import "fmt"

func main() {
	word1 := "zzzyyy"
	word2 := "babababab"

	fmt.Println(checkAlmostEquivalent(word1, word2))
}

func checkAlmostEquivalent(word1 string, word2 string) bool {

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, c := range word1 {
		if _, ok := m1[c]; ok {
			m1[c] += 1
		} else {
			m1[c] = 1
		}
	}

	for _, c := range word2 {
		if _, ok := m2[c]; ok {
			m2[c] += 1
		} else {
			m2[c] = 1
		}
	}

	for k, v := range m1 {
		v2 := m2[k]

		if abs(v-v2) > 3 {
			return false
		}
	}

	for k, v := range m2 {
		v2 := m1[k]

		if abs(v-v2) > 3 {
			return false
		}
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
