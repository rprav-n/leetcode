// Author 		: Praveen
// Date   		: 11/04/2023
// Question 	: https://leetcode.com/problems/toeplitz-matrix/
// Submission 	: https://leetcode.com/problems/toeplitz-matrix/submissions/931680017/

package main

import "fmt"

func main() {
	matrix := [][]int{{36, 59, 71, 15, 26, 82, 87}, {56, 36, 59, 71, 15, 26, 82}, {15, 0, 36, 59, 71, 15, 26}}
	matrix = [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
}

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
