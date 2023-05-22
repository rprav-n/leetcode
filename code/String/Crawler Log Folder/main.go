// Author 		: Praveen
// Date   		: 22/05/2023
// Question 	: https://leetcode.com/problems/crawler-log-folder/
// Submission 	: https://leetcode.com/problems/crawler-log-folder/submissions/955249660/

package main

import (
	"fmt"
)

func main() {
	logs := []string{"./", "wz4/", "../", "mj2/", "../", "../", "ik0/", "il7/"}
	fmt.Println(minOperations(logs))
}

func minOperations(logs []string) int {
	var count int

	for _, s := range logs {
		if s == "../" {
			if count > 0 {
				count -= 1
			}

		} else if s == "./" {

		} else {
			count++
		}
	}

	if count < 0 {
		count = 0
	}

	return count
}
