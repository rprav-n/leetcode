// Author 		: Praveen
// Date   		: 11/01/2023
// Question 	: https://leetcode.com/problems/sign-of-the-product-of-an-array/
// Submission 	: https://leetcode.com/problems/sign-of-the-product-of-an-array/submissions/876223724/

/*
	Question:
	--------------------------------
	There is a function signFunc(x) that returns:

	1 if x is positive.
	-1 if x is negative.
	0 if x is equal to 0.
	You are given an integer array nums. Let product be the product of all values in the array nums.

	Return signFunc(product).


	Example 1:
	Input: nums = [-1,-2,-3,-4,3,2,1]
	Output: 1
	Explanation: The product of all values in the array is 144, and signFunc(144) = 1


*/

package main

import "fmt"

func main() {
	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	fmt.Println(arraySign(nums))
}

func arraySign(nums []int) int {
	var ans int
	p := 1

	for _, n := range nums {
		p = p * n
		if p > 0 {
			p = 1
		} else if p < 0 {
			p = -1
		}
	}

	if p > 0 {
		ans = 1
	} else if p < 0 {
		ans = -1
	} else {
		ans = 0
	}

	return ans
}
