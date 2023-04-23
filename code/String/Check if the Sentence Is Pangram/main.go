// Author 		: Praveen
// Date   		: 23/04/2023
// Question 	: https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// Submission 	: https://leetcode.com/problems/check-if-the-sentence-is-pangram/submissions/938526477/

package main

import "fmt"

func main() {
	sentence := "leetcode"
	fmt.Println(checkIfPangram(sentence))
}

func checkIfPangram(sentence string) bool {

	m := make(map[rune]struct{})

	for _, c := range sentence {
		m[c] = struct{}{}
		if len(m) == 26 {
			return true
		}

	}

	return false
}
