// Author 		: Praveen
// Date   		: 26/12/2022
// Question 	: https://leetcode.com/problems/matrix-diagonal-sum/
// Submission 	: https://leetcode.com/problems/matrix-diagonal-sum/submissions/865777218/

/*
	Question:
	--------------------------------
	Given a square matrix mat, return the sum of the matrix diagonals.

	Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.


	Example 1:
	Input: mat = [[1,2,3],
              	  [4,5,6],
              	  [7,8,9]]
	Output: 25
	Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
	Notice that element mat[1][1] = 5 is counted only once.



*/

package main

import "fmt"

func main() {
	mat := [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	fmt.Println(diagonalSum(mat))
}

func diagonalSum(mat [][]int) int {
	var result int

	l := len(mat[0])

	x := 0
	y := l - 1

	for i := 0; i < l; i++ {
		arr := mat[i]
		if x != y {
			result += arr[x] + arr[y]
		} else {
			result += arr[x]
		}

		x++
		y--
	}

	return result
}
