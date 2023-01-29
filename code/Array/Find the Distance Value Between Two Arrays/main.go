// Author 		: Praveen
// Date   		: 29/01/2023
// Question 	: https://leetcode.com/problems/find-the-distance-value-between-two-arrays/
// Submission 	: https://leetcode.com/problems/find-the-distance-value-between-two-arrays/submissions/887561294/

package main

import "fmt"

func main() {
	arr1 := []int{2, 1, 100, 3}
	arr2 := []int{-5, -2, 10, -3, 7}
	d := 6
	fmt.Println(findTheDistanceValue(arr1, arr2, d))
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var count int

	cl := len(arr2)

	for _, n := range arr1 {
		c := 0
	inner:
		for _, m := range arr2 {
			if Abs(n-m) <= d {
				break inner
			}
			c++
		}
		if c == cl {
			count++
		}
	}

	return count
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
