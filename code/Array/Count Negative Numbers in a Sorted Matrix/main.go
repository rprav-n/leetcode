// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
// Submission 	: https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/submissions/866861016/

/*
	Question:
	--------------------------------
	Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.


	Example 1:
	Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
	Output: 8
	Explanation: There are 8 negatives number in the matrix.


*/

package main

import "fmt"

func main() {
	grid := [][]int{{3, 2}, {-1, -1}}
	fmt.Println(countNegatives(grid))
}

func countNegatives(grid [][]int) int {
	var count int

	l := len(grid)
	for i := 0; i < l; i++ {
		g := grid[i]
		jl := len(g)
		for j := jl - 1; j >= 0; j-- {
			if g[j] < 0 {
				count++
			} else {
				break
			}
		}
	}

	return count
}
