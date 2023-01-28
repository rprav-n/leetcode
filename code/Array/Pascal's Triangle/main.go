// Author 		: Praveen
// Date   		: 28/01/2023
// Question 	: https://leetcode.com/problems/pascals-triangle/
// Submission 	: https://leetcode.com/problems/pascals-triangle/submissions/886845296/

package main

import "fmt"

func main() {
	numRows := 3
	fmt.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	arr := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		arr[i-1] = make([]int, i)
		arr[i-1][0] = 1
		arr[i-1][len(arr[i-1])-1] = 1

		if i > 2 {

			j := 0
			k := 1
			for j < len(arr[i-2]) && k < len(arr[i-2]) {
				n1 := arr[i-2][j]
				n2 := arr[i-2][k]
				arr[i-1][k] = n1 + n2
				j++
				k++
			}

		}
	}

	return arr
}
