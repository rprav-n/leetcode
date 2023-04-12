// Author 		: Praveen
// Date   		: 12/04/2023
// Question 	: https://leetcode.com/problems/count-distinct-numbers-on-board/
// Submission 	: https://leetcode.com/problems/count-distinct-numbers-on-board/submissions/932619025/

package main

import "fmt"

func main() {
	n := 5
	fmt.Println(distinctIntegers(n))
}

func distinctIntegers(n int) int {
	// Hack
	if n == 1 {
		return n
	}
	return n - 1
}
