// Author 		: Praveen
// Date   		: 07/01/2023
// Question 	: https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/
// Submission 	: https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/submissions/873467942/

/*
	Question:
	--------------------------------
	Given an array of integers nums, you start with an initial positive value startValue.

	In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).

	Return the minimum positive value of startValue such that the step by step sum is never less than 1.


	Example 1:
	Input: nums = [-3,2,-3,4,2]
	Output: 5
	Explanation: If you choose startValue = 4, in the third iteration your step by step sum is less than 1.
	step by step sum
	startValue = 4 | startValue = 5 | nums
	(4 -3 ) = 1  | (5 -3 ) = 2    |  -3
	(1 +2 ) = 3  | (2 +2 ) = 4    |   2
	(3 -3 ) = 0  | (4 -3 ) = 1    |  -3
	(0 +4 ) = 4  | (1 +4 ) = 5    |   4
	(4 +2 ) = 6  | (5 +2 ) = 7    |   2


*/

package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, -5, -1}
	fmt.Println(minStartValue(nums))
}

func minStartValue(nums []int) int {
	var ans int

	startValue := 1
	l := len(nums)

	for true {
		c := 0
		f := startValue
	inner:
		for _, n := range nums {
			f = f + n
			if f > 0 {
				c++
			} else {
				break inner
			}
			if c == l {
				return startValue
			}
		}
		startValue++
	}

	return ans
}
