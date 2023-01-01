// Author 		: Praveen
// Date   		: 01/01/2023
// Question 	: https://leetcode.com/problems/kth-distinct-string-in-an-array/
// Submission 	: https://leetcode.com/problems/kth-distinct-string-in-an-array/submissions/869117491/

/*
	Question:
	--------------------------------
	A distinct string is a string that is present only once in an array.

	Given an array of strings arr, and an integer k, return the kth distinct string present in arr. If there are fewer than k distinct strings, return an empty string "".

	Note that the strings are considered in the order in which they appear in the array.


	Example 1:
	Input: arr = ["d","b","c","b","c","a"], k = 2
	Output: "a"
	Explanation:
	The only distinct strings in arr are "d" and "a".
	"d" appears 1st, so it is the 1st distinct string.
	"a" appears 2nd, so it is the 2nd distinct string.
	Since k == 2, "a" is returned.


*/

package main

import (
	"fmt"
)

func main() {
	arr := []string{"b", "a", "c", "a"}
	k := 2
	fmt.Println(kthDistinct(arr, k))
}

func kthDistinct(arr []string, k int) string {
	var result string

	m := make(map[string]int)

	for _, s := range arr {
		if _, ok := m[s]; ok {
			m[s] += 1
		} else {
			m[s] = 1
		}

	}
	fmt.Println(m)
	count := 0
	for _, s := range arr {
		fmt.Println(s, m[s])
		if m[s] == 1 {
			count += 1
		}
		if count == k {
			result = s
			break
		}
	}
	return result
}
