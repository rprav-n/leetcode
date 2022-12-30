// Author 		: Praveen
// Date   		: 30/12/2022
// Question 	: https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order/
// Submission 	: https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order/submissions/868192083/

/*
	Question:
	--------------------------------
	Given the array nums, obtain a subsequence of the array whose sum of elements is strictly greater than the sum of the non included elements in such subsequence.

	If there are multiple solutions, return the subsequence with minimum size and if there still exist multiple solutions, return the subsequence with the maximum total sum of all its elements. A subsequence of an array can be obtained by erasing some (possibly zero) elements from the array.

	Note that the solution with the given constraints is guaranteed to be unique. Also return the answer sorted in non-increasing order.


	Example 1:
	Input: nums = [4,3,10,9,8]
	Output: [10,9]
	Explanation: The subsequences [10,9] and [10,8] are minimal such that the sum of their elements is strictly greater than the sum of elements not included. However, the subsequence [10,9] has the maximum total sum of its elements.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 10, 9, 8}
	fmt.Println(minSubsequence(nums))
}

func minSubsequence(nums []int) []int {
	var result []int

	l := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < l+1; i++ {
		if SumOf(nums[:i]) > SumOf(nums[i:]) {
			result = append(result, nums[:i]...)
			break
		}
	}

	return result
}

func SumOf(arr []int) int {
	var sum int

	for _, n := range arr {
		sum += n
	}

	return sum
}
