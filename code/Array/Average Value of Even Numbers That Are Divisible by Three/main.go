// Author 		: Praveen
// Date   		: 03/04/2023
// Question 	: https://leetcode.com/problems/average-value-of-even-numbers-that-are-divisible-by-three/
// Submission 	: https://leetcode.com/problems/average-value-of-even-numbers-that-are-divisible-by-three/submissions/927360853/

package main

import "fmt"

func main() {
	nums := []int{9, 3, 8, 4, 2, 5, 3, 8, 6, 1}
	fmt.Println(averageValue(nums))
}

func averageValue(nums []int) int {
	var res int
	c := 0
	sum := 0
	for _, n := range nums {
		if n%2 == 0 && n%3 == 0 {
			fmt.Println(n)
			sum += n
			c++
		}
	}
	if c == 0 {
		c = 1
	}
	res = sum / c
	return res
}
