// Author 		: Praveen
// Date   		: 03/04/2023
// Question 	: https://leetcode.com/problems/max-consecutive-ones/
// Submission 	: https://leetcode.com/problems/max-consecutive-ones/submissions/927355983/

package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	var res int

	count := 0

	for _, n := range nums {
		if n == 1 {
			count++
			res = max(count, res)
		} else {
			count = 0
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
