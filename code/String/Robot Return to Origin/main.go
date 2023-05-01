// Author 		: Praveen
// Date   		: 01/05/2023
// Question 	: https://leetcode.com/problems/robot-return-to-origin/
// Submission 	: https://leetcode.com/problems/robot-return-to-origin/submissions/942351829/

package main

import "fmt"

func main() {
	moves := "LL"
	fmt.Println(judgeCircle(moves))
}

func judgeCircle(moves string) bool {

	m := make(map[rune]int)

	for _, c := range moves {
		if _, ok := m[c]; ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}

	if m['U']-m['D'] != 0 || m['L']-m['R'] != 0 {
		return false
	}

	return true
}
