// Author 		: Praveen
// Date   		: 09/01/2023
// Question 	: https://leetcode.com/problems/two-furthest-houses-with-different-colors/
// Submission 	: https://leetcode.com/problems/two-furthest-houses-with-different-colors/submissions/874817308/

/*
	Question:
	--------------------------------
	There are n houses evenly lined up on the street, and each house is beautifully painted. You are given a 0-indexed integer array colors of length n, where colors[i] represents the color of the ith house.

	Return the maximum distance between two houses with different colors.

	The distance between the ith and jth houses is abs(i - j), where abs(x) is the absolute value of x.


	Example 1:
	Input: colors = [1,1,1,6,1,1,1]
	Output: 3
	Explanation: In the above image, color 1 is blue, and color 6 is red.
	The furthest two houses with different colors are house 0 and house 3.
	House 0 has color 1, and house 3 has color 6. The distance between them is abs(0 - 3) = 3.
	Note that houses 3 and 6 can also produce the optimal answer.


*/

package main

import "fmt"

func main() {
	colors := []int{0, 1}
	fmt.Println(maxDistance(colors))
}

func maxDistance(colors []int) int {
	var ans int

	var arr []int

	l := len(colors)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if colors[i] != colors[j] {
				arr = append(arr, Abs(i-j))
			}
		}
	}
	ans = Max(arr)
	return ans
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(arr []int) int {
	var max int

	for _, n := range arr {
		if n > max {
			max = n
		}
	}

	return max
}
