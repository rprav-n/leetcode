// Author 		: Praveen
// Date   		: 26/04/2023
// Question 	: https://leetcode.com/problems/replace-all-digits-with-characters/
// Submission 	: https://leetcode.com/problems/replace-all-digits-with-characters/submissions/940193520/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "a1c1e1"
	fmt.Println(replaceDigits(s))
}

var alphabets = "abcdefghijklmnopqrstuvwxyz"

func replaceDigits(s string) string {
	var ans string

	l := len(s)

	for i := 1; i < l; i += 2 {
		// shift(s[i-1], s[i])
		ch := rune(s[i-1])
		n, _ := strconv.Atoi(string(s[i]))
		ans += rshift(ch, n)
	}
	if l%2 != 0 {
		ans += string(s[l-1])
	}

	return ans
}

func rshift(b rune, i int) string {
	s := string(b)
	index := 0
	for ic, c := range alphabets {
		if c == b {
			index = i + ic
			break
		}
	}
	s += string(alphabets[index])

	return s
}
