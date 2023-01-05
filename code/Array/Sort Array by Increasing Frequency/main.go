// Author 		: Praveen
// Date   		: 05/01/2023
// Question 	: https://leetcode.com/problems/sort-array-by-increasing-frequency/
// Submission 	: https://leetcode.com/problems/sort-array-by-increasing-frequency/submissions/871639494/

/*
	Question:
	--------------------------------
	Given an array of integers nums, sort the array in increasing order based on the frequency of the values. If multiple values have the same frequency, sort them in decreasing order.

	Return the sorted array.


	Example 1:
	Input: nums = [1,1,2,2,2,3]
	Output: [3,1,1,2,2,2]
	Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2, 2, 3}
	fmt.Println(frequencySort(nums))
}

func frequencySort(nums []int) []int {
	var ans []int
	var newNums [][]int
	m := make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n] = m[n] + 1
		} else {
			m[n] = 1
		}
	}

	for k, v := range m {
		newNums = append(newNums, []int{k, v})
	}

	sort.Slice(newNums, func(i, j int) bool {
		condition := newNums[i][1] < newNums[j][1]

		if newNums[i][1] == newNums[j][1] {
			condition = newNums[i][0] > newNums[j][0]
		}

		return condition
	})

	for _, n := range newNums {
		for i := 0; i < n[1]; i++ {
			ans = append(ans, n[0])
		}
	}

	return ans
}
