// Author 		: Praveen
// Date   		: 01/02/2023
// Question 	: https://leetcode.com/problems/shortest-completing-word/description/
// Submission 	: https://leetcode.com/problems/shortest-completing-word/submissions/889548098/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	licensePlate := "1s3 PSt"
	words := []string{"step", "steps", "stripe", "stepple"}
	fmt.Println(shortestCompletingWord(licensePlate, words))
}

func shortestCompletingWord(licensePlate string, words []string) string {
	var ans string

	m := map[rune]int{}

	for _, r := range licensePlate {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			if val, ok := m[r]; ok {
				m[r] = val + 1
			} else {
				m[r] = 1
			}
		}
	}

	correctAns := []string{}

	for _, word := range words {
		mc := make(map[rune]int)

		for k, v := range m {
			mc[k] = v
		}
		for _, r := range word {
			r = unicode.ToLower(r)
			if _, ok := mc[r]; ok {

				mc[r] -= 1

				if mc[r] == 0 {
					delete(mc, r)
				}
			}

		}

		if len(mc) == 0 {
			correctAns = append(correctAns, word)
		}
	}

	minLen := len(correctAns[0])
	for _, word := range correctAns {
		if len(word) < minLen {
			minLen = len(word)
		}
	}

	for _, word := range correctAns {
		if len(word) == minLen {
			ans = word
			break
		}
	}

	return ans
}
