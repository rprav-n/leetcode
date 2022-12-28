// Author 		: Praveen
// Date   		: 28/12/2022
// Question 	: https://leetcode.com/problems/height-checker/
// Submission 	: https://leetcode.com/problems/height-checker/submissions/866864270/

/*
	Question:
	--------------------------------
	A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line in non-decreasing order by height. Let this ordering be represented by the integer array expected where expected[i] is the expected height of the ith student in line.

	You are given an integer array heights representing the current order that the students are standing in. Each heights[i] is the height of the ith student in line (0-indexed).

	Return the number of indices where heights[i] != expected[i].


	Example 1:
	Input: heights = [1,1,4,2,1,3]
	Output: 3
	Explanation:
	heights:  [1,1,4,2,1,3]
	expected: [1,1,1,2,3,4]
	Indices 2, 4, and 5 do not match.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	heights := []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(heights))
}

func heightChecker(heights []int) int {
	var count int
	l := len(heights)
	cpy := make([]int, l)

	copy(cpy, heights)

	sort.Ints(heights)

	for i, h := range heights {
		if h != cpy[i] {
			count++
		}
	}

	return count
}
