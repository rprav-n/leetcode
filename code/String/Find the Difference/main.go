// Author 		: Praveen
// Date   		: 10/06/2023
// Question 	: https://leetcode.com/problems/find-the-difference/
// Submission 	: https://leetcode.com/problems/find-the-difference/submissions/968250150/

package main

import "fmt"

func main() {
	s := "a"
	t := "aa"
	fmt.Println(string(findTheDifference(s, t)))
}

func findTheDifference(s string, t string) byte {
	var ans byte

	m1 := make(map[byte]int)

	l := len(s)

	for i := 0; i < l; i++ {
		if _, ok := m1[s[i]]; ok {
			m1[s[i]] += 1
		} else {
			m1[s[i]] = 1
		}
	}

	l = len(t)

	for i := 0; i < l; i++ {
		b := t[i]
		if val, ok := m1[b]; !ok {
			return b
		} else {
			if val > 0 {
				m1[b] -= 1
			} else {
				return b
			}

		}
	}

	return ans
}
