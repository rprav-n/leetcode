// Author 		: Praveen
// Date   		: 27/01/2023
// Question 	: https://leetcode.com/problems/sort-array-by-parity-ii/
// Submission 	: https://leetcode.com/problems/sort-array-by-parity-ii/submissions/886415320/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 2, 1}
	//			  0, 1, 2, 3
	fmt.Println(sortArrayByParityII(nums))
}

func sortArrayByParityII(nums []int) []int {
	var ans []int

	l := len(nums)
	half := l / 2

	sort.Slice(nums, func(i, j int) bool {
		if nums[i]%2 == 0 {
			return false
		}
		return true
	})

	for i := 0; i < half; i++ {
		ans = append(ans, nums[i+half])
		ans = append(ans, nums[i])
	}

	return ans
}
