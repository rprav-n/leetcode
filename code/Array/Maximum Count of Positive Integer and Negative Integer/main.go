// Author 		: Praveen
// Date   		: 26/01/2023
// Question 	: https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/
// Submission 	: https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/submissions/885531977/

package main

import "fmt"

func main() {
	nums := []int{5, 20, 66, 1314}
	fmt.Println(maximumCount(nums))
}

func maximumCount(nums []int) int {
	var ans int

	max := 0
	min := 0

	for _, n := range nums {
		if n < 0 {
			min++
		} else if n > 0 {
			max++
		}
	}

	if max > min {
		ans = max
	} else {
		ans = min
	}

	return ans
}
