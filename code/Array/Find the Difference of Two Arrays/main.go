// Author 		: Praveen
// Date   		: 03/01/2023
// Question 	: https://leetcode.com/problems/find-the-difference-of-two-arrays/
// Submission 	: https://leetcode.com/problems/find-the-difference-of-two-arrays/submissions/870589980/

/*
	Question:
	--------------------------------
	Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

	answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
	answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
	Note that the integers in the lists may be returned in any order.


	Example 1:
	Input: nums1 = [1,2,3], nums2 = [2,4,6]
	Output: [[1,3],[4,6]]
	Explanation:
	For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
	For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].


*/

package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	fmt.Println(findDifference(nums1, nums2))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	answer := make([][]int, 2)

	for _, n := range nums1 {
		if !Contains(nums2, n) && !Contains(answer[0], n) {
			answer[0] = append(answer[0], n)
		}
	}

	for _, n := range nums2 {
		if !Contains(nums1, n) && !Contains(answer[1], n) {
			answer[1] = append(answer[1], n)
		}
	}

	return answer
}

func Contains(arr []int, val int) bool {
	for _, n := range arr {
		if n == val {
			return true
		}
	}
	return false
}
