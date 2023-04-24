// Author 		: Praveen
// Date   		: 24/04/2023
// Question 	: https://leetcode.com/problems/count-asterisks/
// Submission 	: https://leetcode.com/problems/count-asterisks/submissions/938737095/

package main

import "fmt"

func main() {
	s := "yo|uar|e**|b|e***au|tifu|l"
	fmt.Println(countAsterisks(s))
}

func countAsterisks(s string) int {
	var count int
	pair := 0
	for _, c := range s {
		if pair%2 == 0 && c == '*' {
			count++
		}
		if c == '|' {
			pair++
		}
	}

	return count
}
