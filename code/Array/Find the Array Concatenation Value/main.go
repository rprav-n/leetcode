// Author 		: Praveen
// Date   		: 07/04/2023
// Question 	: https://leetcode.com/problems/find-the-array-concatenation-value/
// Submission 	: https://leetcode.com/problems/find-the-array-concatenation-value/submissions/929461888/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{5, 14, 13, 8, 12}
	fmt.Println(findTheArrayConcVal(nums))
}

func findTheArrayConcVal(nums []int) int64 {
	var res int64
	l := len(nums)
	i := 0
	j := l - 1
	for i <= j {
		if i == j {
			n := nums[i]
			str := fmt.Sprintf("%d", n)
			i, _ := strconv.ParseInt(str, 10, 64)
			res += i
		} else {
			n1 := nums[i]
			n2 := nums[j]
			str := fmt.Sprintf("%d%d", n1, n2)
			i, _ := strconv.ParseInt(str, 10, 64)
			res += i
		}
		i++
		j--
	}

	return res
}
