// Author 		: Praveen
// Date   		: 10/05/2023
// Question 	: https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
// Submission 	: https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/submissions/947938030/

package main

import (
	"fmt"
)

func main() {
	s := "leet"
	fmt.Println(greatestLetter(s))
}

func greatestLetter(s string) string {

	arr := []rune{}

	cm := make(map[rune]struct{})
	sm := make(map[rune]struct{})
	for _, c := range s {
		if c >= 65 && c <= 90 {
			cm[c] = struct{}{}
		} else {
			sm[c] = struct{}{}
		}
	}

	max := 0
	for k := range cm {
		ch := k + 32
		if _, ok := sm[ch]; ok {
			arr = append(arr, k)
			if int(k) > max {
				max = int(k)
			}
		}
	}
	if max == 0 {
		return ""
	}
	return string(rune(max))
}
