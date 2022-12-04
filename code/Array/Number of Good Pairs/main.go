// Author 		: Praveen
// Date   		: 04/12/2022
// Question 	: https://leetcode.com/problems/number-of-good-pairs/
// Submission 	: https://leetcode.com/problems/number-of-good-pairs/submissions/854562534/

/*
	Question:
	--------------------------------

	Given an array of integers nums, return the number of good pairs.

	A pair (i, j) is called good if nums[i] == nums[j] and i < j.


	Example 1:
	Input: nums = [1,2,3,1,1,3]
	Output: 4
	Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

	Example 2:
	Input: nums = [1,1,1,1]
	Output: 6
	Explanation: Each pair in the array are good.

*/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(nums))
}

func numIdenticalPairs(nums []int) int {
	var numOfGoodPairs int
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] && i < j {
				numOfGoodPairs++
			}
		}
	}

	return numOfGoodPairs

}
