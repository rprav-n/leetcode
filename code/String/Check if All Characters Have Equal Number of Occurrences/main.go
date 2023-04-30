// Author 		: Praveen
// Date   		: 30/04/2023
// Question 	: https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
// Submission 	: https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/submissions/942032041/

package main

import "fmt"

func main() {
	s := "aaabb"
	fmt.Println(areOccurrencesEqual(s))
}

func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)

	for _, c := range s {
		if val, ok := m[c]; ok {
			m[c] = val + 1
		} else {
			m[c] = 1
		}
	}

	count := m[rune(s[0])]

	for _, v := range m {
		if v != count {
			return false
		}
	}
	return true
}
