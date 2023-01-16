// Author 		: Praveen
// Date   		: 16/01/2023
// Question 	: https://leetcode.com/problems/missing-number/
// Submission 	: https://leetcode.com/problems/missing-number/submissions/879284543/

/*
	Question:
	--------------------------------
	Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.


	Example 1:
	Input: nums = [3,0,1]
	Output: 2
	Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	var ans int
	sort.Ints(nums)

	if nums[0] != 0 {
		return 0
	}

	l := len(nums)

	for i := 1; i < l; i++ {
		old := nums[i-1]
		if old+1 != nums[i] {
			ans = i
			return ans
		}
	}

	ans = nums[l-1] + 1

	return ans
}
