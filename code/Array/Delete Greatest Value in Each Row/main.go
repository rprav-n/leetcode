// Author 		: Praveen
// Date   		: 23/12/2022
// Question 	: https://leetcode.com/problems/delete-greatest-value-in-each-row/
// Submission 	: https://leetcode.com/problems/delete-greatest-value-in-each-row/submissions/864379648/

/*
	Question:
	--------------------------------
	You are given an m x n matrix grid consisting of positive integers.

	Perform the following operation until grid becomes empty:

	Delete the element with the greatest value from each row. If multiple such elements exist, delete any of them.
	Add the maximum of deleted elements to the answer.
	Note that the number of columns decreases by one after each operation.

	Return the answer after performing the operations described above.


	Example 1:
	Input: grid = [[1,2,4],[3,3,1]]
	Output: 8
	Explanation: The diagram above shows the removed values in each step.
	- In the first operation, we remove 4 from the first row and 3 from the second row (notice that, there are two cells with value 3 and we can remove any of them). We add 4 to the answer.
	- In the second operation, we remove 2 from the first row and 3 from the second row. We add 3 to the answer.
	- In the third operation, we remove 1 from the first row and 1 from the second row. We add 1 to the answer.
	The final answer = 4 + 3 + 1 = 8.


*/

package main

import "fmt"

func main() {
	grid := [][]int{{10}}
	fmt.Println(deleteGreatestValue(grid))
}

func deleteGreatestValue(grid [][]int) int {
	var result int

	pl := len(grid)
	cl := len(grid[0])
	var arr []int
	x := 0
	for x < cl {
		var newArr []int
		i := 0
		for i < pl {
			arr = grid[i]
			newArr = append(newArr, getMaxNumber(grid[i]))
			grid[i] = deleteFromArr(arr)
			i++
		}

		result += getMaxNumber(newArr)
		x++
	}

	return result
}

func deleteFromArr(arr []int) []int {
	var newArr []int
	index := getMaxNumberIndex(arr)
	newArr = remove(arr, index)
	return newArr
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func getMaxNumberIndex(arr []int) int {
	var index, max int

	for i, n := range arr {
		if n > max {
			max = n
			index = i
		}
	}
	return index
}

func getMaxNumber(arr []int) int {
	var max int
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}
