// Author 		: Praveen
// Date   		: https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
// Question 	: 30/12/2022
// Submission 	: https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/submissions/868142951/

/*
	Question:
	--------------------------------
	You are given a non-negative integer array nums. In one operation, you must:

	Choose a positive integer x such that x is less than or equal to the smallest non-zero element in nums.
	Subtract x from every positive element in nums.
	Return the minimum number of operations to make every element in nums equal to 0.


	Example 1:
	Input: nums = [1,5,0,3,5]
	Output: 3
	Explanation:
	In the first operation, choose x = 1. Now, nums = [0,4,0,2,4].
	In the second operation, choose x = 2. Now, nums = [0,2,0,0,2].
	In the third operation, choose x = 2. Now, nums = [0,0,0,0,0].

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 4, 2, 5}
	fmt.Println(minimumOperations(nums))
}

func minimumOperations(nums []int) int {
	var count int

	sort.Ints(nums)

	l := len(nums)

	for i, n := range nums {
		if n == 0 {
			continue
		}
		for j := i; j < l; j++ {
			nums[j] = nums[j] - n
		}
		count++
	}

	return count
}
