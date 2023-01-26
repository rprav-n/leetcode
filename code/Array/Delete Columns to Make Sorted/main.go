// Author 		: Praveen
// Date   		: 26/01/2023
// Question 	: https://leetcode.com/problems/delete-columns-to-make-sorted/
// Submission 	: https://leetcode.com/problems/delete-columns-to-make-sorted/submissions/885554339/

package main

import (
	"fmt"
)

func main() {
	strs := []string{"rrjk", "furt", "guzm"}
	fmt.Println(minDeletionSize(strs))
}

func minDeletionSize(strs []string) int {
	var ans int

	l := len(strs)
	ls := len(strs[0])

	for i := 0; i < ls; i++ {
		newStr := ""
		for _, s := range strs {
			newStr += string(s[i])
		}
		fmt.Println(newStr)
		start := rune(newStr[0])
		for j := 1; j < l; j++ {
			r := rune(newStr[j])
			if start > r {
				ans++
				break
			}
			start = r
		}
	}

	return ans
}
