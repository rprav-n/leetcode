// Author 		: Praveen
// Date   		: 30/12/2022
// Question 	: https://leetcode.com/problems/longest-subsequence-with-limited-sum/
// Submission 	: https://leetcode.com/problems/longest-subsequence-with-limited-sum/submissions/868149007/

/*
	Question:
	--------------------------------
	You are given an integer array nums of length n, and an integer array queries of length m.

	Return an array answer of length m where answer[i] is the maximum size of a subsequence that you can take from nums such that the sum of its elements is less than or equal to queries[i].

	A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.


	Example 1:



*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 4, 5}
	queries := []int{1}
	fmt.Println(answerQueries(nums, queries))
}

func answerQueries(nums []int, queries []int) []int {
	var result []int

	sort.Ints(nums)

	for _, q := range queries {
		var s, c int
		for _, n := range nums {
			s += n
			if s > q {
				break
			}
			c++
		}
		result = append(result, c)
	}

	return result
}
