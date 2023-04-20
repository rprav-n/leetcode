// Author 		: Praveen
// Date   		: 20/04/2023
// Question 	: https://leetcode.com/problems/split-a-string-in-balanced-strings/
// Submission 	: https://leetcode.com/problems/split-a-string-in-balanced-strings/submissions/936700900/

package main

import "fmt"

func main() {
	s := "RLRRLLRLRL"
	fmt.Println(balancedStringSplit(s))
}

func balancedStringSplit(s string) int {
	var count int

	l := len(s)
	cr := 0
	cl := 0
	for i := 0; i < l; i++ {
		c := s[i]

		if c == 'R' {
			cr++
			if cr != 0 && cl != 0 && cr == cl {
				count++
				cl = 0
				cr = 0
			}
			// cl = 0
		} else {
			cl++
			if cr != 0 && cl != 0 && cr == cl {
				count++
				cl = 0
				cr = 0
			}
			// cr = 0
		}

	}

	return count
}
