// Author 		: Praveen
// Date   		: 03/01/2023
// Question 	: https://leetcode.com/problems/minimum-absolute-difference/
// Submission 	: https://leetcode.com/problems/minimum-absolute-difference/submissions/870613128/

/*
	Question:
	--------------------------------
	Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.

	Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows

	a, b are from arr
	a < b
	b - a equals to the minimum absolute difference of any two elements in arr


	Example 1:
	Input: arr = [4,2,1,3]
	Output: [[1,2],[2,3],[3,4]]
	Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{40, 11, 26, 27, -20}
	fmt.Println(minimumAbsDifference(arr))
}

func minimumAbsDifference(arr []int) [][]int {
	var answer [][]int
	l := len(arr)

	sort.Ints(arr)

	minDiff := Abs(arr[0] - arr[1])

	for i := 0; i < l-1; i++ {
		minDiff = Min(minDiff, Abs(arr[i]-arr[i+1]))
	}

	for i := 0; i < l-1; i++ {
		if Abs(arr[i]-arr[i+1]) == minDiff {
			answer = append(answer, []int{arr[i], arr[i+1]})
		}
	}

	return answer
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
