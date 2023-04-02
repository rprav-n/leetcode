// Author 		: Praveen
// Date   		: 02/04/2023
// Question 	: https://leetcode.com/problems/rank-transform-of-an-array/
// Submission 	: https://leetcode.com/problems/rank-transform-of-an-array/submissions/926720324/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{40, 10, 20, 30}
	fmt.Println(arrayRankTransform(arr))
}

func arrayRankTransform(arr []int) []int {
	var res []int
	l := len(arr)
	if l == 0 {
		return res
	}
	m := map[int]int{}
	unsortedArr := make([]int, len(arr))
	copy(unsortedArr, arr)
	sort.Ints(arr)
	rank := 1

	for i := 0; i < l-1; i++ {
		if arr[i] != arr[i+1] {
			m[arr[i]] = rank
			rank++
		}
	}
	m[arr[l-1]] = rank

	for _, n := range unsortedArr {
		res = append(res, m[n])
	}

	return res
}
