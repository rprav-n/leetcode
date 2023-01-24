// Author 		: Praveen
// Date   		: 25/01/2023
// Question 	: https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
// Submission 	: https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/submissions/884572986/

package main

import "fmt"

func main() {
	m := 2
	n := 2
	indices := [][]int{{1, 1}, {0, 0}}
	fmt.Println(oddCells(m, n, indices))
}

func oddCells(m int, n int, indices [][]int) int {
	var ans int
	matrix := make([][]int, m)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for _, index := range indices {
		ri := index[0]
		ci := index[1]
		for i := 0; i < m; i++ {
			if i == ri {
				for j := 0; j < len(matrix[i]); j++ {
					matrix[i][j] += 1
					if matrix[i][j]%2 == 0 {
						ans--
					} else {
						ans++
					}
				}
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i == ci {
					matrix[j][i] += 1
					if matrix[j][i]%2 == 0 {
						ans--
					} else {
						ans++
					}
				}
			}
		}
	}

	return ans
}
