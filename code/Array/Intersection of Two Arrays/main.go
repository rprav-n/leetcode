// Author 		: Praveen
// Date   		: 02/01/2023
// Question 	: https://leetcode.com/problems/intersection-of-two-arrays/
// Submission 	: https://leetcode.com/problems/intersection-of-two-arrays/submissions/869847510/

/*
	Question:
	--------------------------------
	Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.


	Example 1:
	Input: nums1 = [1,2,2,1], nums2 = [2,2]
	Output: [2]


*/

package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	var result []int

	m := make(map[int]struct{})
	l := make(map[int]struct{})

	for _, n := range nums1 {
		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
		}
	}

	for _, n := range nums2 {
		if _, ok := l[n]; !ok {
			l[n] = struct{}{}
		}
	}

	for key := range m {
		if _, ok := l[key]; ok {
			result = append(result, key)
		}
	}

	return result
}
