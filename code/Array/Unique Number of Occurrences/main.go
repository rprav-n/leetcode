// Author 		: Praveen
// Date   		: 28/12/2022
// Question 	: https://leetcode.com/problems/unique-number-of-occurrences/
// Submission 	: https://leetcode.com/problems/unique-number-of-occurrences/submissions/867042950/

/*
	Question:
	--------------------------------
	Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.


	Example 1:
	Input: arr = [1,2,2,1,1,3]
	Output: true
	Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.


*/

package main

import "fmt"

func main() {
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	isUnique := true
	m := make(map[int]int)
	n := make(map[int]struct{})
	for _, n := range arr {
		if _, ok := m[n]; ok {
			m[n] = m[n] + 1
		} else {
			m[n] = 1
		}
	}
	for _, v := range m {
		if _, ok := n[v]; ok {
			isUnique = false
			break
		} else {
			n[v] = struct{}{}
		}
	}
	return isUnique
}
