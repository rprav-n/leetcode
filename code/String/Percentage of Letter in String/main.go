// Author 		: Praveen
// Date   		: 02/05/2023
// Question 	: https://leetcode.com/problems/percentage-of-letter-in-string/
// Submission 	: https://leetcode.com/problems/percentage-of-letter-in-string/submissions/942986787/

package main

import (
	"fmt"
)

func main() {
	s := "vmvvvvvzrvvpvdvvvvyfvdvvvvpkvvbvvkvvfkvvvkvbvvnvvomvzvvvdvvvkvvvvvvvvvlvcvilaqvvhoevvlmvhvkvtgwfvvzy"
	letter := byte('v')

	fmt.Println(percentageLetter(s, letter))
}

func percentageLetter(s string, letter byte) int {

	l := len(s)
	c := 0
	for i := 0; i < l; i++ {
		if s[i] == letter {
			c++
		}
	}
	fmt.Println((float64(c) / float64(l)) * 100)
	fmt.Println((float32(c) / float32(l)) * 100)

	return int((float64(c) / float64(l)) * 100)
}
