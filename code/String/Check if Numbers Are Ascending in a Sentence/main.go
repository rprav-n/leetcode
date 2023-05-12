// Author 		: Praveen
// Date   		: 12/05/2023
// Question 	: https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/
// Submission 	: https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/submissions/949098011/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "hello world 5 x 5"
	fmt.Println(areNumbersAscending(s))
}

func areNumbersAscending(s string) bool {

	words := strings.Split(s, " ")
	prevN := 0
	for _, word := range words {
		n, err := strconv.Atoi(word)
		if err != nil {
			continue
		}
		if n <= prevN {
			return false
		}
		prevN = n
	}

	return true
}
