// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
// Submission 	: https://leetcode.com/problems/find-numbers-with-even-number-of-digits/submissions/866277748/

/*
	Question:
	--------------------------------
	Given an array nums of integers, return how many of them contain an even number of digits.


	Example 1:
	Input: nums = [12,345,2,6,7896]
	Output: 2
	Explanation:
	12 contains 2 digits (even number of digits).
	345 contains 3 digits (odd number of digits).
	2 contains 1 digit (odd number of digits).
	6 contains 1 digit (odd number of digits).
	7896 contains 4 digits (even number of digits).
	Therefore only 12 and 7896 contain an even number of digits.


*/

package main

import "fmt"

func main() {
	nums := []int{555, 901, 482, 1771}
	fmt.Println(findNumbers(nums))
}

func findNumbers(nums []int) int {
	var count int

	for _, num := range nums {
		n := num
		var c int
		for n > 0 {
			c++
			n = n / 10
		}
		if c%2 == 0 {
			count++
		}
	}

	return count
}
