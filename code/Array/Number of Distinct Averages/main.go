// Author 		: Praveen
// Date   		: 05/02/2023
// Question 	: https://leetcode.com/problems/number-of-distinct-averages/
// Submission 	: https://leetcode.com/problems/number-of-distinct-averages/submissions/892175152/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 100}
	fmt.Println(distinctAverages(nums))
}

func distinctAverages(nums []int) int {
	var ans int
	m := map[float32]int{}

	sort.Ints(nums)

	i := 0
	j := len(nums) - 1

	for i < j {

		if val, ok := m[float32(nums[i]+nums[j])/2]; ok {
			m[float32(nums[i]+nums[j])/2] = val + 1
		} else {
			m[float32(nums[i]+nums[j])/2] = 1
		}

		i++
		j--
	}

	ans = len(m)

	return ans
}
