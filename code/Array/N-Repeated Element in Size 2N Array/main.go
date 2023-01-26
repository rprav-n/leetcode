// Author 		: Praveen
// Date   		: 26/01/2023
// Question 	: https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
// Submission 	: https://leetcode.com/problems/n-repeated-element-in-size-2n-array/submissions/885548912/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes(nums))
}

func repeatedNTimes(nums []int) int {
	var ans int

	l := len(nums)
	lhalf := l / 2

	sort.Ints(nums)

	i := 1

	count := 1
	for i < l {
		n1 := nums[i-1]
		n2 := nums[i]
		if n1 == n2 {
			count++
		} else {
			count = 1
		}
		if count == lhalf {
			ans = n1
			break
		}
		i++
	}

	return ans
}
