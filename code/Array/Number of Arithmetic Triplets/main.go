// Author 		: Praveen
// Date   		: 13/12/2022
// Question 	: https://leetcode.com/problems/number-of-arithmetic-triplets/
// Submission 	: https://leetcode.com/problems/number-of-arithmetic-triplets/submissions/859298745/

/*
	Question:
	--------------------------------
	You are given a 0-indexed, strictly increasing integer array nums and a positive integer diff. A triplet (i, j, k) is an arithmetic triplet if the following conditions are met:

	i < j < k,
	nums[j] - nums[i] == diff, and
	nums[k] - nums[j] == diff.
	Return the number of unique arithmetic triplets.


	Example 1:

	Input: nums = [0,1,4,6,7,10], diff = 3
	Output: 2
	Explanation:
	(1, 2, 4) is an arithmetic triplet because both 7 - 4 == 3 and 4 - 1 == 3.
	(2, 4, 5) is an arithmetic triplet because both 10 - 7 == 3 and 7 - 4 == 3.

*/

package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 8, 9}
	diff := 2
	fmt.Println(arithmeticTriplets(nums, diff))
}

func arithmeticTriplets(nums []int, diff int) int {
	var count int
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if i < j && j < k && nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					count++
				}
			}
		}
	}

	return count
}
