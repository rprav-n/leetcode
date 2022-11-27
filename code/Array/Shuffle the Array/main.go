// Author 		: Praveen
// Date   		: 27/11/2022
// Question 	: https://leetcode.com/problems/shuffle-the-array/
// Submission 	: https://leetcode.com/problems/shuffle-the-array/submissions/850745816/

/*
	Question:
	--------------------------------

	Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

	Return the array in the form [x1,y1,x2,y2,...,xn,yn].

	Example 1:
	Input: nums = [2,5,1,3,4,7], n = 3
	Output: [2,3,5,4,1,7]
	Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7].

	Example 2:
	Input: nums = [1,2,3,4,4,3,2,1], n = 4
	Output: [1,4,2,3,3,2,4,1]


	Notes:
	--------------------------------
	Loop through range 0 to < n and append to result arr with 'i' element and 'i+n' element from nums array

*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n := 4
	fmt.Println(shuffle(nums, n))
}

func shuffle(nums []int, n int) []int {
	var result []int

	for i := 0; i < n; i++ {
		result = append(result, []int{nums[i], nums[i+n]}...)
	}

	return result
}
