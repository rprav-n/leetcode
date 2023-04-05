// Author 		: Praveen
// Date   		: 05/04/2023
// Question 	: https://leetcode.com/problems/monotonic-array/
// Submission 	: https://leetcode.com/problems/monotonic-array/submissions/928624169/

package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	fmt.Println(isMonotonic(nums))
}

func isMonotonic(nums []int) bool {
	var res bool
	l := len(nums)

	mi := 0
	md := 0

	for i := 0; i < l-1; i++ {
		if nums[i] <= nums[i+1] {
			mi++
		}

		if nums[i] >= nums[i+1] {
			md++
		}
	}

	if mi == l-1 || md == l-1 {
		res = true
	}

	return res
}
