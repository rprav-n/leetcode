// Author 		: Praveen
// Date   		: 03/01/2023
// Question 	: https://leetcode.com/problems/number-of-unequal-triplets-in-array/
// Submission 	: https://leetcode.com/problems/number-of-unequal-triplets-in-array/submissions/870562233/

/*
	Question:
	--------------------------------
	You are given a 0-indexed array of positive integers nums. Find the number of triplets (i, j, k) that meet the following conditions:

	0 <= i < j < k < nums.length
	nums[i], nums[j], and nums[k] are pairwise distinct.
	In other words, nums[i] != nums[j], nums[i] != nums[k], and nums[j] != nums[k].
	Return the number of triplets that meet the conditions.


	Example 1:
	Input: nums = [4,4,2,4,3]
	Output: 3
	Explanation: The following triplets meet the conditions:
	- (0, 2, 4) because 4 != 2 != 3
	- (1, 2, 4) because 4 != 2 != 3
	- (2, 3, 4) because 2 != 4 != 3
	Since there are 3 triplets, we return 3.
	Note that (2, 0, 4) is not a valid triplet because 2 > 0.


*/

package main

import "fmt"

func main() {
	nums := []int{4, 4, 2, 4, 3}
	fmt.Println(unequalTriplets(nums))
}

func unequalTriplets(nums []int) int {
	var count int
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			for k := j; k < l; k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					count++
				}
			}
		}
	}

	return count
}
