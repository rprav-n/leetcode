// Author 		: Praveen
// Date   		: 04/05/2023
// Question 	: https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/
// Submission 	: https://leetcode.com/problems/check-if-number-has-equal-digit-count-and-digit-value/submissions/944159901/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "1210"
	fmt.Println(digitCount(num))
}

func digitCount(num string) bool {

	m := make(map[string]int)

	for _, c := range num {
		s := string(c)
		if val, ok := m[s]; ok {
			m[s] = val + 1
		} else {
			m[s] = 1
		}
	}
	for i, c := range num {
		is := strconv.Itoa(i)
		cs := string(c)
		ci, _ := strconv.Atoi(cs)
		if m[is] != ci {
			return false
		}
	}

	return true
}
