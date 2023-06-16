// Author 		: Praveen
// Date   		: 16/06/2023
// Question 	: https://leetcode.com/problems/ransom-note/
// Submission 	: https://leetcode.com/problems/ransom-note/submissions/972796252/

package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "ba"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	mm := make(map[rune]int)

	for _, c := range magazine {
		if _, ok := mm[c]; ok {
			mm[c] += 1
		} else {
			mm[c] = 1
		}
	}

	for _, c := range ransomNote {
		val := mm[c]

		if val > 0 {
			mm[c] -= 1
		} else {
			return false
		}
	}

	return true
}
