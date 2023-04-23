// Author 		: Praveen
// Date   		: 23/04/2023
// Question 	: https://leetcode.com/problems/merge-strings-alternately/
// Submission 	: https://leetcode.com/problems/merge-strings-alternately/submissions/938530369/

package main

import "fmt"

func main() {
	word1 := "ab"
	word2 := "pqrs"

	fmt.Println(mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	var res string

	l1 := len(word1)
	l2 := len(word2)

	i := 0
	j := 0

	for i < l1 || j < l2 {

		if i < l1 {
			res += string(word1[i])
		}

		if j < l2 {
			res += string(word2[j])
		}
		i++
		j++
	}

	return res
}
