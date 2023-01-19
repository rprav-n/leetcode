// Author 		: Praveen
// Date   		: 19/01/2023
// Question 	: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// Submission 	: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/submissions/881301012/

package main

import "fmt"

func main() {
	nums := []int{1, 1}
	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	var ans []int
	m := make(map[int]struct{})
	l := len(nums)

	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
		}
	}

	for i := 1; i <= l; i++ {
		if _, ok := m[i]; !ok {
			ans = append(ans, i)
		}
	}

	return ans
}
