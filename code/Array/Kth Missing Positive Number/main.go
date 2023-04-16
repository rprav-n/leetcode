// Author 		: Praveen
// Date   		: 16/04/2023
// Question 	: https://leetcode.com/problems/kth-missing-positive-number/
// Submission 	: https://leetcode.com/problems/kth-missing-positive-number/submissions/934813544/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 6, 7}
	k := 2

	fmt.Println(findKthPositive(arr, k))
}

func findKthPositive(arr []int, k int) int {

	startIndex := 0
	start := 1
	l := len(arr)

	for k >= 1 {

		if startIndex >= l || start != arr[startIndex] {
			startIndex--
			k--
		}
		if k == 0 {
			return start
		}
		startIndex++
		start++
	}

	return start
}
