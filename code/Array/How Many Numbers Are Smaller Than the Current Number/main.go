// Author 		: Praveen
// Date   		: 06/12/2022
// Question 	: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
// Submission 	: https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/submissions/855671037/

/*
	Question:
	--------------------------------

	Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it. That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

	Return the answer in an array.

	Example 1:
	Input: nums = [8,1,2,2,3]
	Output: [4,0,1,1,3]
	Explanation:
	For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3).
	For nums[1]=1 does not exist any smaller number than it.
	For nums[2]=2 there exist one smaller number than it (1).
	For nums[3]=2 there exist one smaller number than it (1).
	For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).


	Example 2:
	Input: nums = [6,5,4,8]
	Output: [2,1,0,3]


	Notes:
	--------------------------------
	Result order is Important

	1. Copy the nums to another slice
	2. Sort the nums array
	3. Loop through 1 to len(nums) - 1
		... (In Progress)

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{8, 1, 2, 2, 3}
	nums = []int{6, 5, 3, 8}
	nums = []int{7, 7, 7, 7}
	fmt.Println(smallerNumbersThanCurrent(nums))
}

func smallerNumbersThanCurrent(nums []int) []int {
	var result []int
	oldNums := make([]int, len(nums))
	copy(oldNums, nums)

	resultMap := make(map[int]int)

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 0 {
			resultMap[nums[i]] = i
		} else {
			j := 0
			k := i
			for k > 0 {
				if nums[k]-nums[k-1] != 0 {
					j = k
				}
				k--
			}
			resultMap[nums[j]] = j
		}
	}

	for _, n := range oldNums {
		result = append(result, resultMap[n])
	}

	return result
}
