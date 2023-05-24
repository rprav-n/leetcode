// Author 		: Praveen
// Date   		: 24/05/2023
// Question 	: https://leetcode.com/problems/occurrences-after-bigram/
// Submission 	: https://leetcode.com/problems/occurrences-after-bigram/submissions/956563952/

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "we will we will rock you"
	first := "we"
	second := "will"
	fmt.Println(findOcurrences(text, first, second))
}

func findOcurrences(text string, first string, second string) []string {
	var ans []string

	words := strings.Split(text, " ")

	l := len(words)

	i := 0
	j := 1

	for i < l && j < l {
		if words[i] == first && words[j] == second {
			if j+1 < l {
				ans = append(ans, words[j+1])
			}

		}
		i++
		j++
	}

	return ans
}
