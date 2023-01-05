// Author 		: Praveen
// Date   		: 05/01/2023
// Question 	: https://leetcode.com/problems/find-common-characters/
// Submission 	: https://leetcode.com/problems/find-common-characters/submissions/871998594/

/*
	Question:
	--------------------------------
	Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates). You may return the answer in any order.


	Example 1:
	Input: words = ["bella","label","roller"]
	Output: ["e","l","l"]


*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"cool", "lock", "cook"}
	fmt.Println(commonChars(words))
}

func commonChars(words []string) []string {
	var ans []string
	l := len(words)
	r := l - 1
	f := words[0]

	m := make(map[rune]int)

	for _, ch := range f {
		c := 0
		for i := 1; i < l; i++ {
			word := words[i]

			if strings.ContainsRune(word, ch) {
				c++
				words[i] = NewWord(word, ch)
			}
		}
		if c == r {
			ans = append(ans, string(ch))
			m[ch] = 1
		}
	}

	return ans
}

func NewWord(s string, ch rune) string {
	var word string

	word = strings.Replace(s, string(ch), "", 1)

	return word
}
