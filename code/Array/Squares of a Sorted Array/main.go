// Author 		: Praveen
// Date   		: 31/12/2022
// Question 	: https://leetcode.com/problems/squares-of-a-sorted-array/
// Submission 	: https://leetcode.com/problems/squares-of-a-sorted-array/submissions/868658587/

/*
	Question:
	--------------------------------
	Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.


	Example 1:
	Input: nums = [-4,-1,0,3,10]
	Output: [0,1,9,16,100]
	Explanation: After squaring, the array becomes [16,1,0,9,100].
	After sorting, it becomes [0,1,9,16,100].


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {

	var result []int

	sort.Slice(nums, func(i, j int) bool {
		a := nums[i]
		b := nums[j]
		return a*a < b*b
	})

	for _, n := range nums {
		result = append(result, n*n)
	}

	return result
}
