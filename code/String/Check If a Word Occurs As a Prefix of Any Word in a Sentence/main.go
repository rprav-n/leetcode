// Author 		: Praveen
// Date   		: 23/05/2023
// Question 	: https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
// Submission 	: https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/submissions/955903968/

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "this problem is an easy problem"
	searchWord := "pro"
	fmt.Println(isPrefixOfWord(sentence, searchWord))
}

func isPrefixOfWord(sentence string, searchWord string) int {
	count := -1

	l := len(searchWord)

	words := strings.Split(sentence, " ")

	for i, word := range words {
		if len(word) >= l && word[:l] == searchWord {
			count = i + 1
			break
		}
	}

	return count
}
