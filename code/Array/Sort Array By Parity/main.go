// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/sort-array-by-parity/
// Submission 	: https://leetcode.com/problems/sort-array-by-parity/submissions/866475495/

/*
	Question:
	--------------------------------
	Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.

	Return any array that satisfies this condition.


	Example 1:
	Input: nums = [3,1,2,4]
	Output: [2,4,3,1]
	Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.


*/

package main

import "fmt"

func main() {
	nums := []int{0}
	fmt.Println(sortArrayByParity(nums))
}

func sortArrayByParity(nums []int) []int {
	var result []int
	var odd, even []int
	for _, n := range nums {
		if n%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}

	result = append(result, even...)
	result = append(result, odd...)

	return result
}
