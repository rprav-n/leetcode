// Author 		: Praveen
// Date   		: 19/12/2022
// Question 	: https://leetcode.com/problems/sort-the-people/
// Submission 	: https://leetcode.com/problems/sort-the-people/submissions/862230937/

/*
	Question:
	--------------------------------
	You are given an array of strings names, and an array heights that consists of distinct positive integers. Both arrays are of length n.

	For each index i, names[i] and heights[i] denote the name and height of the ith person.

	Return names sorted in descending order by the people's heights.


	Example 1:
	Input: names = ["Mary","John","Emma"], heights = [180,165,170]
	Output: ["Mary","Emma","John"]
	Explanation: Mary is the tallest, followed by Emma and John.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	fmt.Println(sortPeople(names, heights))
}

func sortPeople(names []string, heights []int) []string {
	var result []string
	m := make(map[int]string)

	l := len(names)

	for i := 0; i < len(names); i++ {
		m[heights[i]] = names[i]
	}

	sort.Ints(heights)

	for i := l - 1; i >= 0; i-- {
		result = append(result, m[heights[i]])
	}

	return result
}
