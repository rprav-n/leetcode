// Author 		: Praveen
// Date   		: 03/01/2022
// Question 	: https://leetcode.com/problems/single-number/
// Submission 	: https://leetcode.com/problems/single-number/submissions/870572400/

/*
	Question:
	--------------------------------
	Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

	You must implement a solution with a linear runtime complexity and use only constant extra space.


	Example 1:
	Input: nums = [2,2,1]
	Output: 1


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	var result int
	l := len(nums)

	sort.Ints(nums)

	i := 0
	j := 1

	for i < l && j < l {
		if nums[i] != nums[j] {
			result = nums[i]
			return result
		}
		i = j + 1
		j = i + 1
	}

	result = nums[l-1]

	return result
}
