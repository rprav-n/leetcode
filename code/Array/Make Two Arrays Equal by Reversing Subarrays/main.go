// Author 		: Praveen
// Date   		: 30/12/2022
// Question 	: https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/
// Submission 	: https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/submissions/868186253/

/*
	Question:
	--------------------------------
	You are given two integer arrays of equal length target and arr. In one step, you can select any non-empty subarray of arr and reverse it. You are allowed to make any number of steps.

	Return true if you can make arr equal to target or false otherwise.


	Example 1:
	Input: target = [1,2,3,4], arr = [2,4,1,3]
	Output: true
	Explanation: You can follow the next steps to convert arr to target:
	1- Reverse subarray [2,4,1], arr becomes [1,4,2,3]
	2- Reverse subarray [4,2], arr becomes [1,2,4,3]
	3- Reverse subarray [4,3], arr becomes [1,2,3,4]
	There are multiple ways to convert arr to target, this is not the only way to do so.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	target := []int{3, 7, 9}
	arr := []int{3, 7, 11}
	fmt.Println(canBeEqual(target, arr))
}

func canBeEqual(target []int, arr []int) bool {

	lt := len(target)
	la := len(arr)
	if lt != la {
		return false
	}

	sort.Ints(target)
	sort.Ints(arr)

	i := 0
	j := lt
	for i < j {
		if target[i] != arr[i] {
			return false
		}
		i++
	}
	return true
}
