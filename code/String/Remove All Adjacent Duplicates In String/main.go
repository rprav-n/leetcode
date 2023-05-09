// Author 		: Praveen
// Date   		: 09/05/2023
// Question 	: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
// Submission 	: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/submissions/947312285/

package main

import "fmt"

func main() {
	s := "azxxzy"
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(s string) string {
	var ans string

	arr := []rune{}

	for _, c := range s {
		la := len(arr)
		if la == 0 {
			arr = append(arr, c)
		} else {
			if arr[la-1] == c {
				arr = arr[:la-1]
			} else {
				arr = append(arr, c)
			}
		}
	}

	l := len(arr)

	for i := 0; i < l; i++ {
		ans += string(arr[i])
	}

	return ans
}
