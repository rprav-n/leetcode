// Author 		: Praveen
// Date   		: 02/05/2023
// Question 	: https://leetcode.com/problems/first-letter-to-appear-twice/
// Submission 	: https://leetcode.com/problems/first-letter-to-appear-twice/submissions/942980707/

package main

import "fmt"

func main() {
	s := "abcdd"
	fmt.Println(repeatedCharacter(s))
}

func repeatedCharacter(s string) byte {
	var ans byte

	l := len(s)
	m := make(map[byte]int)

	for i := 0; i < l; i++ {
		b := s[i]
		if val, ok := m[b]; ok {
			if val == 2 {
				ans = b
				break
			}
			m[b] = val + 1
			if m[b] == 2 {
				ans = b
				break
			}
		} else {
			m[b] = 1
		}
	}

	return ans
}
