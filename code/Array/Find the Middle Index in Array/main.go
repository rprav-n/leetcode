// Author 		: Praveen
// Date   		: 09/01/2023
// Question 	: https://leetcode.com/problems/find-the-middle-index-in-array/
// Submission 	: https://leetcode.com/problems/find-the-middle-index-in-array/submissions/874806284/

/*
	Question:
	--------------------------------
	Given a 0-indexed integer array nums, find the leftmost middleIndex (i.e., the smallest amongst all the possible ones).

	A middleIndex is an index where nums[0] + nums[1] + ... + nums[middleIndex-1] == nums[middleIndex+1] + nums[middleIndex+2] + ... + nums[nums.length-1].

	If middleIndex == 0, the left side sum is considered to be 0. Similarly, if middleIndex == nums.length - 1, the right side sum is considered to be 0.

	Return the leftmost middleIndex that satisfies the condition, or -1 if there is no such index.


	Example 1:
	Input: nums = [2,3,-1,8,4]
	Output: 3
	Explanation: The sum of the numbers before index 3 is: 2 + 3 + -1 = 4
	The sum of the numbers after index 3 is: 4 = 4


*/

package main

import "fmt"

func main() {
	nums := []int{4, 0}
	fmt.Println(findMiddleIndex(nums))
}

func findMiddleIndex(nums []int) int {
	var ans int
	var arr []int
	l := len(nums)

	if l == 1 {
		return 0
	}

	for i := 0; i < l; i++ {
		leftSum := SumOf(nums[0:i])
		rightSum := SumOf(nums[i+1:])

		if leftSum == rightSum {
			arr = append(arr, i)
		}
	}

	if len(arr) == 0 {
		ans = -1
	} else {
		ans = MinOf(arr)
	}

	return ans
}

func SumOf(arr []int) int {
	var sum int

	for _, n := range arr {
		sum += n
	}

	return sum
}

func MinOf(arr []int) int {
	min := arr[0]

	for _, n := range arr {
		if n < min {
			min = n
		}
	}

	return min
}
