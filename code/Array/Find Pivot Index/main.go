// Author 		: Praveen
// Date   		: 06/04/2023
// Question 	: https://leetcode.com/problems/find-pivot-index/
// Submission 	: https://leetcode.com/problems/find-pivot-index/submissions/929178293/

package main

import "fmt"

func main() {
	nums := []int{-1, -1, 0, 0, -1, -1}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	index := -1
	l := len(nums)
	for i := 0; i < l; i++ {
		lSum := 0
		rSum := 0
		lSum += sumOf(nums[:i])
		rSum += sumOf(nums[i+1:])

		if lSum == rSum {
			index = i
			break
		}
	}

	return index
}

func sumOf(arr []int) int {
	sum := 0

	for _, n := range arr {
		sum += n
	}

	return sum
}
