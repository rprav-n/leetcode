// Author 		: Praveen
// Date   		: 04/01/2023
// Question 	: https://leetcode.com/problems/intersection-of-multiple-arrays/
// Submission 	: https://leetcode.com/problems/intersection-of-multiple-arrays/submissions/871278316/

/*
	Question:
	--------------------------------
	Given a 2D integer array nums where nums[i] is a non-empty array of distinct positive integers, return the list of integers that are present in each array of nums sorted in ascending order.


	Example 1:
	Input: nums = [[3,1,2,4,5],[1,2,3,4],[3,4,5,6]]
	Output: [3,4]
	Explanation:
	The only integers present in each of nums[0] = [3,1,2,4,5], nums[1] = [1,2,3,4], and nums[2] = [3,4,5,6] are 3 and 4, so we return [3,4].


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(intersection(nums))
}

func intersection(nums [][]int) []int {
	var ans []int

	l := len(nums)

	arr := nums[0]

	r := l - 1

	for _, n := range arr {
		s := 0
		for i := 1; i < l; i++ {
			if contains(nums[i], n) {
				s++
			}
		}
		if s >= r {
			ans = append(ans, n)
		}
	}

	sort.Ints(ans)

	return ans
}

func contains(arr []int, val int) bool {
	for _, n := range arr {
		if n == val {
			return true
		}
	}
	return false
}
