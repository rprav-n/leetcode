// Author 		: Praveen
// Date   		: 05/05/2023
// Question 	: https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/
// Submission 	: https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/submissions/945059492/

package main

import "fmt"

func main() {
	s := "aababcabc"
	fmt.Println(countGoodSubstrings(s))
}

func countGoodSubstrings(s string) int {
	var count int

	i := 0
	j := 1
	k := 2

	l := len(s)

	for k < l {
		if s[i] != s[j] && s[j] != s[k] && s[k] != s[i] {
			count++
		}
		i++
		j++
		k++
	}

	return count
}
