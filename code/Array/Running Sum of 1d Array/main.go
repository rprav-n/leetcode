// Author 		: Praveen
// Date   		: 25/11/2022
// Question 	: https://leetcode.com/problems/running-sum-of-1d-array
// Submission 	: https://leetcode.com/problems/running-sum-of-1d-array/submissions/849726663/

/*
	Question:
	--------------------------------

	Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

	Return the running sum of nums.

	Example 1:
	Input: nums = [1,2,3,4]
	Output: [1,3,6,10]
	Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4]

	Example 2:
	Input: nums = [1,1,1,1,1]
	Output: [1,2,3,4,5]
	Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1]


	Notes:
	--------------------------------
	Create 2 variable one for result arr ([]int) and for oldSum (int)

	Simple solution, loop through the nums arr and do addition assignment to oldSum with the current loop number and append the
	oldSum to the results arr.

	Then return the result arr.

*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	var result []int
	var oldSum int
	for i := 0; i < len(nums); i++ {
		oldSum += nums[i]
		result = append(result, oldSum)
	}

	return result
}
