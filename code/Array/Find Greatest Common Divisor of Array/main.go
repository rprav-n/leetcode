// Author 		: Praveen
// Date   		: 22/01/2023
// Question 	: https://leetcode.com/problems/find-greatest-common-divisor-of-array/
// Submission 	: https://leetcode.com/problems/find-greatest-common-divisor-of-array/submissions/882613317/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 3}
	fmt.Println(findGCD(nums))
}

func findGCD(nums []int) int {
	var ans int
	l := len(nums)
	sort.Ints(nums)

	min := nums[0]
	max := nums[l-1]

	for i := 1; i <= max; i++ {
		if min%i == 0 && max%i == 0 && i > ans {
			ans = i
		}
	}

	return ans
}
