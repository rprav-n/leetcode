// Author 		: Praveen
// Date   		: 26/12/2022
// Question 	: https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
// Submission 	: https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/submissions/865771686/

/*
	Question:
	--------------------------------
	Given the array of integers nums, you will choose two different indices i and j of that array. Return the maximum value of (nums[i]-1)*(nums[j]-1).


	Example 1:
	Input: nums = [3,4,5,2]
	Output: 12
	Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12.



*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 5, 4, 5}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	return (nums[l-1] - 1) * (nums[l-2] - 1)
}
