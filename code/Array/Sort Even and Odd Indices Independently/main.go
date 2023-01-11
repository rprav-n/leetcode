// Author 		: Praveen
// Date   		: 11/01/2023
// Question 	: https://leetcode.com/problems/sort-even-and-odd-indices-independently/
// Submission 	: https://leetcode.com/problems/sort-even-and-odd-indices-independently/submissions/876211359/

/*
	Question:
	--------------------------------
	You are given a 0-indexed integer array nums. Rearrange the values of nums according to the following rules:

	Sort the values at odd indices of nums in non-increasing order.
	For example, if nums = [4,1,2,3] before this step, it becomes [4,3,2,1] after. The values at odd indices 1 and 3 are sorted in non-increasing order.
	Sort the values at even indices of nums in non-decreasing order.
	For example, if nums = [4,1,2,3] before this step, it becomes [2,1,4,3] after. The values at even indices 0 and 2 are sorted in non-decreasing order.
	Return the array formed after rearranging the values of nums.


	Example 1:
	Input: nums = [4,1,2,3]
	Output: [2,3,4,1]
	Explanation:
	First, we sort the values present at odd indices (1 and 3) in non-increasing order.
	So, nums changes from [4,1,2,3] to [4,3,2,1].
	Next, we sort the values present at even indices (0 and 2) in non-decreasing order.
	So, nums changes from [4,1,2,3] to [2,3,4,1].
	Thus, the array formed after rearranging the values is [2,3,4,1].


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 39, 33, 5, 12, 27, 20, 45, 14, 25, 32, 33, 30, 30, 9, 14, 44, 15, 21}
	fmt.Println(sortEvenOdd(nums))
}

func sortEvenOdd(nums []int) []int {
	var ans []int

	var oddArr, evenArr []int

	for i, n := range nums {
		if i%2 == 0 {
			evenArr = append(evenArr, n)
		} else {
			oddArr = append(oddArr, n)
		}
	}

	sort.Slice(oddArr, func(i, j int) bool {
		return oddArr[i] > oddArr[j]
	})

	sort.Ints(evenArr)

	for i := 0; i < len(nums); i++ {
		if i < len(evenArr) {
			ans = append(ans, evenArr[i])
		}

		if i < len(oddArr) {
			ans = append(ans, oddArr[i])
		}

	}

	return ans
}
