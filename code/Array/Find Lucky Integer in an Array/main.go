// Author 		: Praveen
// Date   		: 14/01/2023
// Question 	: https://leetcode.com/problems/find-lucky-integer-in-an-array/
// Submission 	: https://leetcode.com/problems/find-lucky-integer-in-an-array/submissions/878072792/

/*
	Question:
	--------------------------------
	Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.

	Return the largest lucky integer in the array. If there is no lucky integer return -1.


	Example 1:
	Input: arr = [2,2,3,4]
	Output: 2
	Explanation: The only lucky number in the array is 2 because frequency[2] == 2.


*/

package main

import "fmt"

func main() {
	arr := []int{2, 2, 2, 3, 3}
	fmt.Println(findLucky(arr))
}

func findLucky(arr []int) int {
	ans := -1
	m := make(map[int]int)

	for _, n := range arr {
		if _, ok := m[n]; ok {
			m[n] = m[n] + 1
		} else {
			m[n] = 1
		}
	}

	for k, v := range m {
		if k == v {
			if k > ans {
				ans = k
			}
		}
	}

	return ans
}
