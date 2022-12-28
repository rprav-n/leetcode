// Author 		: Praveen
// Date   		: 28/12/2022
// Question 	: https://leetcode.com/problems/divide-array-into-equal-pairs/
// Submission 	: https://leetcode.com/problems/divide-array-into-equal-pairs/submissions/866866894/

/*
	Question:
	--------------------------------
	You are given an integer array nums consisting of 2 * n integers.

	You need to divide nums into n pairs such that:

	Each element belongs to exactly one pair.
	The elements present in a pair are equal.
	Return true if nums can be divided into n pairs, otherwise return false.


	Example 1:
	Input: nums = [3,2,3,2,2,2]
	Output: true
	Explanation:
	There are 6 elements in nums, so they should be divided into 6 / 2 = 3 pairs.
	If nums is divided into the pairs (2, 2), (3, 3), and (2, 2), it will satisfy all the conditions.


*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(divideArray(nums))
}

func divideArray(nums []int) bool {
	var hasPairs bool
	m := make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] = m[n] + 1
		} else {
			m[n] = 1
		}
	}
	hasPairs = true
	for _, val := range m {
		if val%2 != 0 {
			hasPairs = false
			break
		}
	}

	return hasPairs
}
