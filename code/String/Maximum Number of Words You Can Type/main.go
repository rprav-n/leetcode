// Author 		: Praveen
// Date   		: 08/05/2023
// Question 	: https://leetcode.com/problems/maximum-number-of-words-you-can-type/
// Submission 	: https://leetcode.com/problems/maximum-number-of-words-you-can-type/submissions/946753003/

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "leet code"
	brokenLetters := "e"
	fmt.Println(canBeTypedWords(text, brokenLetters))
}

func canBeTypedWords(text string, brokenLetters string) int {

	words := strings.Split(text, " ")

	count := len(words)

	for _, word := range words {

		for _, c := range brokenLetters {
			if strings.ContainsRune(word, c) {
				count--
				break
			}
		}
	}

	return count
}
