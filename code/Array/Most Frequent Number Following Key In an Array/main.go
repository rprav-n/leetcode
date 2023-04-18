// Author 		: Praveen
// Date   		: 18/04/2023
// Question 	: https://leetcode.com/problems/most-frequent-number-following-key-in-an-array/
// Submission 	: https://leetcode.com/problems/most-frequent-number-following-key-in-an-array/submissions/935575824/

package main

import "fmt"

func main() {
	nums := []int{2, 2, 2, 2, 3}
	key := 2
	fmt.Println(mostFrequent(nums, key))
}

func mostFrequent(nums []int, key int) int {
	var ans int
	m := make(map[int]int)

	l := len(nums)

	for i := 1; i < l; i++ {
		if nums[i-1] == key {
			if _, ok := m[nums[i]]; ok {
				m[nums[i]] += 1
			} else {
				m[nums[i]] = 1
			}
		}
	}
	var maxV, maxK int
	for k, v := range m {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}
	ans = maxK

	return ans
}
