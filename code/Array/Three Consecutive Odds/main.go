// Author 		: Praveen
// Date   		: 15/01/2023
// Question 	: https://leetcode.com/problems/three-consecutive-odds/
// Submission 	: https://leetcode.com/problems/three-consecutive-odds/submissions/878691031/

/*
	Question:
	--------------------------------
	Given an integer array arr, return true if there are three consecutive odd numbers in the array. Otherwise, return false.


	Example 1:
	Input: arr = [2,6,4,1]
	Output: false
	Explanation: There are no three consecutive odds.


*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	fmt.Println(threeConsecutiveOdds(arr))
}

func threeConsecutiveOdds(arr []int) bool {
	var ans bool

	c := 0
	for _, n := range arr {
		if n%2 != 0 {
			c++
		} else {
			c = 0
		}
		if c == 3 {
			ans = true
		}
	}

	return ans
}
