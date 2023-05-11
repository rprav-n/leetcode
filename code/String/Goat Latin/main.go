// Author 		: Praveen
// Date   		: 11/05/2023
// Question 	: https://leetcode.com/problems/goat-latin/
// Submission 	: https://leetcode.com/problems/goat-latin/submissions/948584974/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	sentence := "The quick brown fox jumped over the lazy dog"
	fmt.Println(toGoatLatin(sentence))
}

func toGoatLatin(sentence string) string {
	var res string
	vowels := "aeiou"

	vm := make(map[rune]struct{})

	for _, c := range vowels {
		vm[c] = struct{}{}
	}

	words := strings.Split(sentence, " ")
	l := len(words)

	for i, word := range words {
		ch := rune(word[0])
		ch = unicode.ToLower(ch)
		if _, ok := vm[ch]; ok {
			res += word + "ma"
		} else {
			f := word[0]
			res += word[1:] + string(f) + "ma"
		}

		j := i + 1
		for j > 0 {
			res += "a"
			j--
		}
		if i != l-1 {
			res += " "
		}

	}

	return res
}
