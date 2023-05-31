// Author 		: Praveen
// Date   		: 31/05/2023
// Question 	: https://leetcode.com/problems/sum-of-digits-of-string-after-convert/
// Submission 	: https://leetcode.com/problems/sum-of-digits-of-string-after-convert/submissions/961130042/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "leetcode"
	k := 2

	fmt.Println(getLucky(s, k))
}

func getLucky(s string, k int) int {
	var sum int

	alphabets := "abcdefghijklmnopqrstuvwxyz"
	m := make(map[rune]int)

	for i, c := range alphabets {
		m[c] = i + 1
	}
	count := 0

	newS := ""
	for _, c := range s {
		newS += fmt.Sprintf("%d", m[c])
	}

	for count < k {
		sum = 0
		for _, c := range newS {
			in, _ := strconv.Atoi(string(c))
			sum += in
		}

		newS = fmt.Sprintf("%d", sum)

		count++
	}

	return sum
}
