// Author 		: Praveen
// Date   		: 14/03/2023
// Question 	: https://leetcode.com/problems/count-special-quadruplets/
// Submission 	: https://leetcode.com/problems/count-special-quadruplets/submissions/933433026/

package main

import "fmt"

func main() {
	nums := []int{9, 6, 8, 23, 39, 23}
	fmt.Println(countQuadruplets(nums))
}

func countQuadruplets(nums []int) int {
	var res int

	m := make(map[int]int)

	for i, n := range nums {
		m[i] = n
	}
	l := len(nums)
	for i := 0; i < l; i++ {
		a := nums[i]
		for j := 0; j < l; j++ {
			b := nums[j]
			if i >= j {
				continue
			}
			for k := 0; k < l; k++ {
				c := nums[k]
				if i >= k || j >= k {
					continue
				}
				sum := a + b + c
				for index, val := range m {
					if val == sum && k < index {
						res++
					}
				}

			}
		}
	}

	return res
}
