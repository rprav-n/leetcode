// Author 		: Praveen
// Date   		: 09/04/2023
// Question 	: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/
// Submission 	: https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/submissions/930613070/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := [][]int{{1, 2}, {2, 3}, {4, 5}}
	nums2 := [][]int{{1, 4}, {3, 2}, {4, 1}}
	fmt.Println(mergeArrays(nums1, nums2))
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var ans [][]int
	m := make(map[int]int)

	ans = append(ans, nums1...)
	ans = append(ans, nums2...)
	sort.Slice(ans, func(i, j int) bool {
		id1 := ans[i][0]
		id2 := ans[j][0]
		return id1 < id2
	})
	res := make([][]int, 0, len(m))
	for _, arr := range ans {
		m[arr[0]] += arr[1]
	}

	for k, v := range m {
		res = append(res, []int{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		id1 := res[i][0]
		id2 := res[j][0]
		return id1 < id2
	})

	return res
}
