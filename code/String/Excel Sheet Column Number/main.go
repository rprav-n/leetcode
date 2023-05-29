// Author 		: Praveen
// Date   		: 29/05/2023
// Question 	: https://leetcode.com/problems/excel-sheet-column-number/
// Submission 	: https://leetcode.com/problems/excel-sheet-column-number/submissions/959759951/

package main

import "fmt"

func main() {
	columnTitle := "CAB"
	fmt.Println(titleToNumber(columnTitle))
}

func titleToNumber(columnTitle string) int {

	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	m := make(map[rune]int)

	for i, c := range alphabets {
		m[c] = i + 1
	}

	l := len(columnTitle)

	oldAns := 0
	for _, c := range columnTitle {
		ans := 1
		lc := l - 1
		val := m[c]
		ans *= val
		for lc > 0 {
			ans *= 26
			lc--
		}
		oldAns += ans
		l--

	}

	return oldAns
}
