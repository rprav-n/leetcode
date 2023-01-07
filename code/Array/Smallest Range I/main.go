// Author 		: Praveen
// Date   		: 07/01/2023
// Question 	: https://leetcode.com/problems/smallest-range-i/
// Submission 	: https://leetcode.com/problems/smallest-range-i/submissions/873455334/

/*
	Question:
	--------------------------------
	You are given an integer array nums and an integer k.

	In one operation, you can choose any index i where 0 <= i < nums.length and change nums[i] to nums[i] + x where x is an integer from the range [-k, k]. You can apply this operation at most once for each index i.

	The score of nums is the difference between the maximum and minimum elements in nums.

	Return the minimum score of nums after applying the mentioned operation at most once for each index in it.


	Example 1:
	Input: nums = [1], k = 0
	Output: 0
	Explanation: The score is max(nums) - min(nums) = 1 - 1 = 0.

	Example 2:
	Input: nums = [0,10], k = 2
	Output: 6
	Explanation: Change nums to be [2, 8]. The score is max(nums) - min(nums) = 8 - 2 = 6.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 6}
	k := 3
	fmt.Println(smallestRangeI(nums, k))
}

func smallestRangeI(nums []int, k int) int {
	var score int

	sort.Ints(nums)

	for i := range nums {
		if i == 0 {
			nums[i] = nums[i] + k
		} else {
			if nums[i] < nums[i-1] {
				diff := nums[i-1] - nums[i]
				if diff >= -k && diff <= k {
					nums[i] = nums[i] + diff
				} else {
					nums[i] = nums[i] + k
				}
			} else {
				diff := nums[i] - nums[i-1]
				if diff >= -k && diff <= k {
					nums[i] = nums[i] - diff
				} else {
					nums[i] = nums[i] - k
				}
			}
		}
	}

	score = Max(nums) - Min(nums)

	return score
}

func Max(arr []int) int {
	max := 0
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}

func Min(arr []int) int {
	min := arr[0]
	for _, n := range arr {
		if n < min {
			min = n
		}
	}

	return min

}
