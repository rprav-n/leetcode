// Author 		: Praveen
// Date   		: 13/04/2023
// Question 	: https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements/
// Submission 	: https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements/submissions/933215387/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3, 3, 3, 3}
	fmt.Println(countElements(nums))
}

func countElements(nums []int) int {
	var count int

	sort.Ints(nums)

	l := len(nums)

	for i := 1; i < l-1; i++ {
		num := nums[i]
		left := containsLess(nums[:i], num)
		right := containsGreater(nums[i+1:], num)
		if left && right {
			count++
		}
	}

	return count
}

func containsLess(arr []int, num int) bool {
	for _, n := range arr {
		if n < num {
			return true
		}
	}

	return false
}
func containsGreater(arr []int, num int) bool {
	for _, n := range arr {
		if n > num {
			return true
		}
	}

	return false
}
