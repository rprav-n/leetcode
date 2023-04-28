// Author 		: Praveen
// Date   		: 28/04/2023
// Question 	: https://leetcode.com/problems/determine-if-string-halves-are-alike/
// Submission 	: https://leetcode.com/problems/determine-if-string-halves-are-alike/submissions/941224827/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "textbook"
	fmt.Println(halvesAreAlike(s))
}

func halvesAreAlike(s string) bool {

	l := len(s)
	h := l / 2
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	var lc, rc int

	for i, c := range s {

		ch := unicode.ToLower(c)
		if _, ok := vowels[ch]; ok {
			if i < h {
				lc++
			} else {
				rc++
				if rc > lc {
					return false
				}
			}
		}

	}
	if lc == rc {
		return true
	}

	return false
}
