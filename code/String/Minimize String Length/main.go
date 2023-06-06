// Author 		: Praveen
// Date   		: 06/06/2023
// Question 	: https://leetcode.com/problems/minimize-string-length/
// Submission 	: https://leetcode.com/problems/minimize-string-length/submissions/965344934/

package main

import "fmt"

func main() {
	s := "aaabbcdaa"
	// fmt.Println(s[:2])
	fmt.Println(minimizedStringLength(s))
}

func minimizedStringLength(s string) int {
	m := make(map[rune]struct{})

	for _, c := range s {
		m[c] = struct{}{}
	}

	return len(m)
}
