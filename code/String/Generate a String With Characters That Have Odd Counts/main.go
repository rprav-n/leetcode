// Author 		: Praveen
// Date   		: 29/04/2023
// Question 	: https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/
// Submission 	: https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/submissions/941716162/

package main

import "fmt"

func main() {
	n := 2
	fmt.Println(generateTheString(n))
}

func generateTheString(n int) string {
	var ans string

	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			ans += "a"
		}
		ans += "b"
	} else {
		for i := 0; i < n; i++ {
			ans += "a"
		}
	}

	return ans
}
