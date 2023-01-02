// Author 		: Praveen
// Date   		: 02/01/2023
// Question 	: https://leetcode.com/problems/smallest-index-with-equal-value/
// Submission 	: https://leetcode.com/problems/smallest-index-with-equal-value/submissions/869404172/

/*
	Question:
	--------------------------------
	Given a 0-indexed integer array nums, return the smallest index i of nums such that i mod 10 == nums[i], or -1 if such index does not exist.

	x mod y denotes the remainder when x is divided by y.


	Example 1:
	Input: nums = [0,1,2]
	Output: 0
	Explanation:
	i=0: 0 mod 10 = 0 == nums[0].
	i=1: 1 mod 10 = 1 == nums[1].
	i=2: 2 mod 10 = 2 == nums[2].
	All indices have i mod 10 == nums[i], so we return the smallest index 0.


*/

package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 1}
	fmt.Println(smallestEqual(nums))
}

func smallestEqual(nums []int) int {
	ans := -1

	for i, n := range nums {
		if i%10 == n {
			ans = i
			break
		}
	}

	return ans
}
