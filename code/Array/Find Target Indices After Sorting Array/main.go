// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/find-target-indices-after-sorting-array/
// Submission 	: https://leetcode.com/problems/find-target-indices-after-sorting-array/submissions/866462298/

/*
	Question:
	--------------------------------
	You are given a 0-indexed integer array nums and a target element target.

	A target index is an index i such that nums[i] == target.

	Return a list of the target indices of nums after sorting nums in non-decreasing order. If there are no target indices, return an empty list. The returned list must be sorted in increasing order.


	Example 1:
	Input: nums = [1,2,5,2,3], target = 2
	Output: [1,2]
	Explanation: After sorting, nums is [1,2,2,3,5].
	The indices where nums[i] == 2 are 1 and 2.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 5, 2, 3}
	target := 5
	fmt.Println(targetIndices(nums, target))
}

func targetIndices(nums []int, target int) []int {
	var result []int

	sort.Ints(nums)

	for i, n := range nums {
		if n == target {
			result = append(result, i)
		}
	}

	return result
}
