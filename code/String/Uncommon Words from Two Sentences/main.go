// Author 		: Praveen
// Date   		: 13/05/2023
// Question 	: https://leetcode.com/problems/uncommon-words-from-two-sentences/
// Submission 	: https://leetcode.com/problems/uncommon-words-from-two-sentences/submissions/949744688/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	fmt.Println(uncommonFromSentences(s1, s2))
}

func uncommonFromSentences(s1 string, s2 string) []string {
	var ans []string

	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")

	m := make(map[string]int)

	for _, word := range words1 {
		if _, ok := m[word]; ok {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}

	for _, word := range words2 {
		if _, ok := m[word]; ok {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}

	for k, v := range m {
		if v == 1 {
			ans = append(ans, k)
		}
	}

	return ans
}
