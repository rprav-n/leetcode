// Author 		: Praveen
// Date   		: 26/12/2022
// Question 	: https://leetcode.com/problems/find-the-highest-altitude/
// Submission 	: https://leetcode.com/problems/find-the-highest-altitude/submissions/865866868/

/*
	Question:
	--------------------------------
	There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. The biker starts his trip on point 0 with altitude equal 0.

	You are given an integer array gain of length n where gain[i] is the net gain in altitude between points i​​​​​​ and i + 1 for all (0 <= i < n). Return the highest altitude of a point.


	Example 1:
	Input: gain = [-5,1,5,0,-7]
	Output: 1
	Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.

	Example 2:
	Input: gain = [-4,-3,-2,-1,4,3,2]
	Output: 0
	Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.


*/

package main

import "fmt"

func main() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}

func largestAltitude(gain []int) int {
	var high, start int

	for i := 0; i < len(gain); i++ {
		start = start + gain[i]
		if start > high {
			high = start
		}
	}

	return high
}
