// Author 		: Praveen
// Date   		: 01/05/2023
// Question 	: https://leetcode.com/problems/remove-palindromic-subsequences/
// Submission 	: https://leetcode.com/problems/remove-palindromic-subsequences/submissions/942350262/

package main

import "fmt"

func main() {
	s := "baabb"
	fmt.Println(removePalindromeSub(s))
}

func removePalindromeSub(s string) int {
	count := 1

	l := len(s)

	i := 0
	j := l - 1

	for i < j {
		if s[i] != s[j] {
			count = 2
		}
		i++
		j--
	}

	return count
}
