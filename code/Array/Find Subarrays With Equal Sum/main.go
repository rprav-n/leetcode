// Author 		: Praveen
// Date   		: 14/03/2023
// Question 	: https://leetcode.com/problems/find-subarrays-with-equal-sum/
// Submission 	: https://leetcode.com/problems/find-subarrays-with-equal-sum/submissions/933425757/

package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0}
	//0, 1, 2, 3, 4
	fmt.Println(findSubarrays(nums))
}

func findSubarrays(nums []int) bool {
	var res bool

	m := make(map[int]string)

	l := len(nums)

	for i := 0; i < l-1; i++ {
		n1 := nums[i]
		n2 := nums[i+1]
		sum := n1 + n2
		str := fmt.Sprintf("%d_%d", i, i+1)
		if val, ok := m[sum]; ok && str != val && str != Reverse(val) {
			return true
		}
		m[sum] = str
	}

	return res
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
