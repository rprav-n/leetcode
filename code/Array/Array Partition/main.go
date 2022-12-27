// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/array-partition/
// Submission 	: https://leetcode.com/problems/array-partition/submissions/866450127/

/*
	Question:
	--------------------------------
	Given an integer array nums of 2n integers, group these integers into n pairs (a1, b1), (a2, b2), ..., (an, bn) such that the sum of min(ai, bi) for all i is maximized. Return the maximized sum.


	Example 1:
	Input: nums = [1,4,3,2]
	Output: 4
	Explanation: All possible pairings (ignoring the ordering of elements) are:
	1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
	2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
	3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
	So the maximum possible sum is 4.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	var result int

	sort.Ints(nums)

	l := len(nums)

	i := 0
	j := 1
	for i < l && j < l {
		result += min(nums[i], nums[j])
		i += 2
		j += 2
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
