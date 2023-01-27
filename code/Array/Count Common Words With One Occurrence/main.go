// Author 		: Praveen
// Date   		: 27/01/2023
// Question 	: https://leetcode.com/problems/count-common-words-with-one-occurrence/
// Submission 	: https://leetcode.com/problems/count-common-words-with-one-occurrence/submissions/886419481/

package main

import "fmt"

func main() {
	words1 := []string{"b", "bb", "bbb"}
	words2 := []string{"a", "aa", "aaa"}
	fmt.Println(countWords(words1, words2))
}

func countWords(words1 []string, words2 []string) int {
	var count int

	m1 := map[string]int{}
	m2 := map[string]int{}

	for _, word := range words1 {
		m1[word] = m1[word] + 1
	}

	for _, word := range words2 {
		m2[word] = m2[word] + 1
	}

	if len(m1) > len(m2) {
		for k, v := range m1 {
			if m2[k] == 1 && v == 1 {
				count++
			}
		}
	} else {
		for k, v := range m2 {
			if m1[k] == 1 && v == 1 {
				count++
			}
		}
	}

	return count
}
