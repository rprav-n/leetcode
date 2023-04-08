// Author 		: Praveen
// Date   		: 08/04/2023
// Question 	: https://leetcode.com/problems/separate-the-digits-in-an-array/
// Submission 	: https://leetcode.com/problems/separate-the-digits-in-an-array/submissions/930061823/

package main

import (
	"fmt"
)

func main() {
	nums := []int{13, 25, 83, 77}
	fmt.Println(separateDigits(nums))
}

func separateDigits(nums []int) []int {
	var ans []int

	for _, n := range nums {
		i := n
		arr := []int{}
		for i > 0 {
			v := i % 10
			i = i / 10
			arr = append(arr, v)
		}

		for j := len(arr) - 1; j >= 0; j-- {
			ans = append(ans, arr[j])
		}
	}

	return ans
}
