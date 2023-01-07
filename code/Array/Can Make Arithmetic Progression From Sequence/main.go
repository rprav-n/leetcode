// Author 		: Praveen
// Date   		: 07/01/2023
// Question 	: https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
// Submission 	: https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/submissions/873253717/

/*
	Question:
	--------------------------------
	A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.

	Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.


	Example 1:
	Input: arr = [3,5,1]
	Output: true
	Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.


*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{0, 0, 1}
	fmt.Println(canMakeArithmeticProgression(arr))
}

func canMakeArithmeticProgression(arr []int) bool {
	l := len(arr)

	sort.Ints(arr)
	fmt.Println(arr)

	i := 0
	j := 1

	diff := math.MaxInt

	for i < l && j < l {
		sub := Abs(arr[i] - arr[j])
		fmt.Println(arr[i], arr[j], sub)
		if i != 0 && sub != diff {
			return false
		}
		diff = sub
		i = i + 1
		j = j + 1
	}

	return true
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
