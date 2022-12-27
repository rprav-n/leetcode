// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/maximum-number-of-pairs-in-array/
// Submission 	: https://leetcode.com/problems/maximum-number-of-pairs-in-array/submissions/866459334/

/*
	Question:
	--------------------------------
	You are given a 0-indexed integer array nums. In one operation, you may do the following:

	Choose two integers in nums that are equal.
	Remove both integers from nums, forming a pair.
	The operation is done on nums as many times as possible.

	Return a 0-indexed integer array answer of size 2 where answer[0] is the number of pairs that are formed and answer[1] is the number of leftover integers in nums after doing the operation as many times as possible.


	Example 1:
	Input: nums = [1,3,2,1,3,2,2]
	Output: [3,1]
	Explanation:
	Form a pair with nums[0] and nums[3] and remove them from nums. Now, nums = [3,2,3,2,2].
	Form a pair with nums[0] and nums[2] and remove them from nums. Now, nums = [2,2,2].
	Form a pair with nums[0] and nums[1] and remove them from nums. Now, nums = [2].
	No more pairs can be formed. A total of 3 pairs have been formed, and there is 1 number leftover in nums.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(numberOfPairs(nums))
}

func numberOfPairs(nums []int) []int {
	var arr []int
	var pair, leftover int

	l := len(nums)

	sort.Ints(nums)

	i := 0
	j := 1

	for i < l && j < l {
		if nums[i] == nums[j] {
			pair++
			nums[i] = -1
			nums[j] = -1
		}

		i++
		j++
	}

	for _, n := range nums {
		if n != -1 {
			leftover++
		}
	}

	arr = append(arr, pair, leftover)

	return arr
}
