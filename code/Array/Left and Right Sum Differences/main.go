// Author 		: Praveen
// Date   		: 04/04/2023
// Question 	: https://leetcode.com/problems/left-and-right-sum-differences/
// Submission 	: https://leetcode.com/problems/left-and-right-sum-differences/submissions/928018823/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1}
	fmt.Println(leftRigthDifference(nums))
}

func leftRigthDifference(nums []int) []int {
	var res []int

	l := len(nums)

	var ls, rs []int

	for i := 0; i < l; i++ {
		ls = append(ls, sumOf(nums[:i]))
		rs = append(rs, sumOf(nums[i+1:]))
	}

	for i, n := range ls {
		res = append(res, abs(n-rs[i]))
	}

	return res
}

func sumOf(arr []int) int {
	s := 0
	for _, n := range arr {
		s += n
	}
	return s
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
