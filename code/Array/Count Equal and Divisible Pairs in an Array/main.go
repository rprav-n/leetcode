// Author 		: Praveen
// Date   		: 26/12/2022
// Question 	: https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/
// Submission 	: https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/submissions/865104473/

/*
	Question:
	--------------------------------
	Given a 0-indexed integer array nums of length n and an integer k, return the number of pairs (i, j) where 0 <= i < j < n, such that nums[i] == nums[j] and (i * j) is divisible by k.


	Example 1:
	Input: nums = [3,1,2,2,2,1,3], k = 2
	Output: 4
	Explanation:
	There are 4 pairs that meet all the requirements:
	- nums[0] == nums[6], and 0 * 6 == 0, which is divisible by 2.
	- nums[2] == nums[3], and 2 * 3 == 6, which is divisible by 2.
	- nums[2] == nums[4], and 2 * 4 == 8, which is divisible by 2.
	- nums[3] == nums[4], and 3 * 4 == 12, which is divisible by 2.



*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	k := 1

	fmt.Println(countPairs(nums, k))
}

func countPairs(nums []int, k int) int {
	var count int
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] && i*j%k == 0 {
				count++
			}
		}
	}

	return count
}
