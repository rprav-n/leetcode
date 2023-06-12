// Author 		: Praveen
// Date   		: 12/06/2023
// Question 	: https://leetcode.com/problems/circular-sentence/
// Submission 	: https://leetcode.com/problems/circular-sentence/submissions/969235721/

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Leetcode is cool"
	fmt.Println(isCircularSentence(sentence))
}

func isCircularSentence(sentence string) bool {

	words := strings.Split(sentence, " ")

	l := len(words)

	for i := 0; i < l; i++ {
		word := words[i]
		if i == l-1 {
			firstWord := words[0]
			if word[len(word)-1] != firstWord[0] {
				return false
			}
		} else {
			secondWord := words[i+1]
			if word[len(word)-1] != secondWord[0] {
				return false
			}
		}
	}

	return true
}
