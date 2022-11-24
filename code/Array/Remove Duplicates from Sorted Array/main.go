// Author 		: Praveen
// Date   		: 24/11/2022
// Question 	: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Submission 	: https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/844671799/

/*
	Question:
	--------------------------------

	Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.

	Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

	Return k after placing the final result in the first k slots of nums.

	Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.


	Example 1:
	Input: nums = [1,1,2]
	Output: 2, nums = [1,2,_]
	Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).

	Notes + Solution:
	--------------------------------
	Datastructure used: map/hashmap (key:number, value:empty struct)

	why empty struct? - Empty struct occupies 0 byte of storage

	Solution: Algorithm is to loop through the nums array and check if the number is present in the map or not
	if NOT -> the number is not duplicate so add it to the map, since we have to modify nums array in-place we're
	updating the nums[result] (which is the total no. of duplicate elements, by default int zero value is zero) to the loop number
	and increment the result variable

*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}

	// Example 2
	// nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	var result int
	numsMap := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := numsMap[n]; !ok {
			numsMap[n] = struct{}{}
			nums[result] = n
			result++
		}
	}
	// fmt.Println(nums)
	return result
}
