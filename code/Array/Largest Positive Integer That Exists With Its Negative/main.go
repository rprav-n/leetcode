// Author 		: Praveen
// Date   		: 08/01/2023
// Question 	: https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/
// Submission 	: https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/submissions/874027430/

/*
	Question:
	--------------------------------
	Given an integer array nums that does not contain any zeros, find the largest positive integer k such that -k also exists in the array.

	Return the positive integer k. If there is no such integer, return -1.


	Example 1:
	Input: nums = [-1,2,-3,3]
	Output: 3
	Explanation: 3 is the only valid k we can find in the array.

	Input: nums = [-1,10,6,7,-7,1]
	Output: 7
	Explanation: Both 1 and 7 have their corresponding negative values in the array. 7 has a larger value.

*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{-10, 8, 6, 7, -2, -3}
	fmt.Println(findMaxK(nums))
}

func findMaxK(nums []int) int {
	var ans int
	var arr []int

	posArr := []int{}
	negArr := []int{}

	for _, n := range nums {
		if n > 0 {
			posArr = append(posArr, n)
		} else {
			negArr = append(negArr, n)
		}
	}

	for _, p := range posArr {
		for _, n := range negArr {
			if p == -n {
				arr = append(arr, p)
			}
		}
	}

	l := len(arr)

	if l == 1 {
		ans = arr[0]
		return ans
	}
	if l == 0 {
		ans = -1
		return ans
	}

	ans = Max(arr)

	return ans
}

func Max(arr []int) int {
	var m int
	for _, n := range arr {
		if n > m {
			m = n
		}
	}

	return m
}
