// Author 		: Praveen
// Date   		: 15/01/2023
// Question 	: https://leetcode.com/problems/maximum-ascending-subarray-sum/
// Submission 	: https://leetcode.com/problems/maximum-ascending-subarray-sum/submissions/878695624/

/*
	Question:
	--------------------------------
	Given an array of positive integers nums, return the maximum possible sum of an ascending subarray in nums.

	A subarray is defined as a contiguous sequence of numbers in an array.

	A subarray [numsl, numsl+1, ..., numsr-1, numsr] is ascending if for all i where l <= i < r, numsi  < numsi+1. Note that a subarray of size 1 is ascending.


	Example 1:
	Input: nums = [10,20,30,5,10,50]
	Output: 65
	Explanation: [5,10,50] is the ascending subarray with the maximum sum of 65.


*/

package main

import "fmt"

func main() {
	nums := []int{12, 17, 15, 13, 10, 11, 12}
	fmt.Println(maxAscendingSum(nums))
}

func maxAscendingSum(nums []int) int {

	nums = append(nums, 0)
	max := 0
	arr := []int{0}
	for _, n := range nums {
		l := len(arr)

		if arr[l-1] < n {
			arr = append(arr, n)
		} else {
			sum := sumOf(arr)
			if sum > max {
				max = sum
			}
			arr = []int{0}
			arr = append(arr, n)
		}

	}
	return max
}

func sumOf(arr []int) int {
	var sum int

	for _, n := range arr {
		sum += n
	}

	return sum
}
