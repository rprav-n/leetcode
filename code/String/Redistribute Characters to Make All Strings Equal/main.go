// Author 		: Praveen
// Date   		: 16/05/2023
// Question 	: https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
// Submission 	: https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/submissions/972790939/

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"caaaaa", "aaaaaaaaa", "a", "bbb", "bbbbbbbbb", "bbb", "cc", "cccccccccccc", "ccccccc", "ccccccc", "cc", "cccc", "c", "cccccccc", "c"}
	fmt.Println(makeEqual(words))
}

func makeEqual(words []string) bool {

	l := len(words)
	if l == 1 {
		return true
	}

	s := strings.Join(words, "")

	m := make(map[rune]int)

	for _, c := range s {
		if val, ok := m[c]; ok {
			m[c] = val + 1
		} else {
			m[c] = 1
		}
	}

	fmt.Println(m)

	for _, v := range m {
		if v%l != 0 {
			return false
		}
	}

	return true
}
