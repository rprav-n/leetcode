// Author 		: Praveen
// Date   		: 23/11/2022
// Question 	: https://leetcode.com/problems/two-sum/
// Submission 	: https://leetcode.com/problems/two-sum/submissions/848673976/

/*
	Question:
	--------------------------------

	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	You can return the answer in any order.

	Example 1:
	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

	Example 2:
	Input: nums = [3,2,4], target = 6
	Output: [1,2]

	Example 3:
	Input: nums = [3,3], target = 6
	Output: [0,1]


	Notes:
	--------------------------------
	Datastructure used: map/hashmap (key:number, value:index)

	Algorithm is to loop through the nums array and check if the target's complement of loop number is present in a map,
	if NOT present - add to map with key as number, value as its index
	if present - return the current loop number index and the map value of this number

	number's complement - if current loop number is 6 and target is 9 -> target-number=complement | 9-6=3
	The nine's complement of 6 is 3

*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	// Example 2
	// nums = []int{3, 2, 4}
	// target = 6

	// Example 3
	// nums = []int{3, 3}
	// target = 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var result []int

	targetsComplementMap := make(map[int]int)

	for i, n := range nums {
		if value, ok := targetsComplementMap[target-n]; ok {
			result = append(result, []int{i, value}...)
			break
		} else {
			targetsComplementMap[n] = i
		}
	}

	return result
}
