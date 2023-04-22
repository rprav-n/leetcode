// Author 		: Praveen
// Date   		: 22/04/2023
// Question 	: https://leetcode.com/problems/sorting-the-sentence/
// Submission 	: https://leetcode.com/problems/sorting-the-sentence/submissions/938010406/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := "is2 sentence4 This1 a3"
	fmt.Println(sortSentence(s))
}

func sortSentence(s string) string {
	var ans string

	arr := [10]string{}

	word := ""
	for _, c := range s {
		if unicode.IsNumber(c) {
			n, _ := strconv.Atoi(string(c))
			arr[n] = word
			word = ""
		} else if c != rune(' ') {
			word += string(c)
		}
	}

	for _, w := range arr {
		if w != "" {
			w = strings.TrimSpace(w)
			ans += w
			ans += " "
		}

	}
	ans = strings.TrimSpace(ans)
	return ans
}
