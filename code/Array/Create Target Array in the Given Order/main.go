// Author 		: Praveen
// Date   		: 09/12/2022
// Question 	: https://leetcode.com/problems/create-target-array-in-the-given-order/
// Submission 	: https://leetcode.com/problems/create-target-array-in-the-given-order/submissions/857230254/

/*
	Question:
	--------------------------------
	Given two arrays of integers nums and index. Your task is to create target array under the following rules:

	Initially target array is empty.
	From left to right read nums[i] and index[i], insert at index index[i] the value nums[i] in target array.
	Repeat the previous step until there are no elements to read in nums and index.
	Return the target array.

	It is guaranteed that the insertion operations will be valid.


	Example 1:
	Input: nums = [0,1,2,3,4], index = [0,1,2,2,1]
	Output: [0,4,1,3,2]
	Explanation:
	nums       index     target
	0            0        [0]
	1            1        [0,1]
	2            2        [0,1,2]
	3            2        [0,1,3,2]
	4            1        [0,4,1,3,2]



	Notes:
	--------------------------------


*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 0}
	index := []int{0, 1, 2, 3, 0}

	fmt.Println(createTargetArray(nums, index))
}

func insert(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func createTargetArray(nums []int, index []int) []int {
	l := len(nums)
	var target []int

	for i := 0; i < l; i++ {
		target = insert(target, index[i], nums[i])
	}

	return target
}
