// Author 		: Praveen
// Date   		: 27/05/2023
// Question 	: https://leetcode.com/problems/valid-anagram/
// Submission 	: https://leetcode.com/problems/valid-anagram/submissions/958510917/

package main

import "fmt"

func main() {
	s := "cat"
	t := "rat"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {

	l := len(s)
	if l != len(t) {
		return false
	}

	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for i := 0; i < l; i++ {
		a := s[i]
		b := t[i]
		if _, ok := m1[a]; ok {
			m1[a] += 1
		} else {
			m1[a] = 1
		}

		if _, ok := m2[b]; ok {
			m2[b] += 1
		} else {
			m2[b] = 1
		}
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
