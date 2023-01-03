// Author 		: Praveen
// Date   		: 03/01/2023
// Question 	: https://leetcode.com/problems/lucky-numbers-in-a-matrix/
// Submission 	: https://leetcode.com/problems/lucky-numbers-in-a-matrix/submissions/870150592/

/*
	Question:
	--------------------------------
	Given an m x n matrix of distinct numbers, return all lucky numbers in the matrix in any order.

	A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.


	Example 1:
	Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
	Output: [15]
	Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column.


*/

package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	fmt.Println(luckyNumbers(matrix))
}

func luckyNumbers(matrix [][]int) []int {
	var ans []int
	m := len(matrix)    // rows
	n := len(matrix[0]) // cols

	for i := 0; i < m; i++ {
		minVal := min(matrix[i])
		for k := 0; k < n; k++ {
			var arr []int
			for l := 0; l < m; l++ {
				arr = append(arr, matrix[l][k])
			}
			maxVal := max(arr)
			if minVal == maxVal {
				ans = append(ans, maxVal)
			}
		}

	}

	return ans
}

func min(arr []int) int {
	var low int
	for _, n := range arr {
		if low == 0 {
			low = n
		}
		if n < low {
			low = n
		}
	}
	return low
}
func max(arr []int) int {
	var high int
	for _, n := range arr {
		if n > high {
			high = n
		}
	}
	return high
}
