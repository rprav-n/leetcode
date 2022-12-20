// Author 		: Praveen
// Date   		: 20/12/2022
// Question 	: https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
// Submission 	: https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/submissions/862451483/

/*
	Question:
	--------------------------------
	Given an integer array nums and an integer k, return the number of pairs (i, j) where i < j such that |nums[i] - nums[j]| == k.

	The value of |x| is defined as:

	x if x >= 0.
	-x if x < 0.


	Example 1:
	Input: nums = [1,2,2,1], k = 1
	Output: 4
	Explanation: The pairs with an absolute difference of 1 are:
	- [1,2,2,1]
	- [1,2,2,1]
	- [1,2,2,1]
	- [1,2,2,1]

	Example 2:
	Input: nums = [1,3], k = 3
	Output: 0
	Explanation: There are no pairs with an absolute difference of 3.


*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	fmt.Println(countKDifference(nums, k))
}

func countKDifference(nums []int, k int) int {
	var count int
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i < j {
				v := absInt(nums[i], nums[j])
				if v == k {
					count++
				}
			}
		}
	}

	return count
}

func absInt(x, y int) int {
	return absDiffInt(x, y)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
