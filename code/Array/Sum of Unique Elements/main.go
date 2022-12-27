// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/sum-of-unique-elements/
// Submission 	: https://leetcode.com/problems/sum-of-unique-elements/submissions/866472656/

/*
	Question:
	--------------------------------
	You are given an integer array nums. The unique elements of an array are the elements that appear exactly once in the array.

	Return the sum of all the unique elements of nums.


	Example 1:
	Input: nums = [1,2,3,2]
	Output: 4
	Explanation: The unique elements are [1,3], and the sum is 4.


*/

package main

import "fmt"

func main() {
	nums := []int{10}
	fmt.Println(sumOfUnique(nums))
}

func sumOfUnique(nums []int) int {

	var sum int
	m := make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] = m[n] + 1
		} else {
			m[n] = 0
		}
	}

	for k, v := range m {
		if v == 0 {
			sum += k
		}
	}
	return sum
}
