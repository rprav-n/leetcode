// Author 		: Praveen
// Date   		: 24/12/2022
// Question 	: https://leetcode.com/problems/count-good-triplets/
// Submission 	: https://leetcode.com/problems/count-good-triplets/submissions/864886618/

/*
	Question:
	--------------------------------
	Given an array of integers arr, and three integers a, b and c. You need to find the number of good triplets.

	A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are true:

	0 <= i < j < k < arr.length
	|arr[i] - arr[j]| <= a
	|arr[j] - arr[k]| <= b
	|arr[i] - arr[k]| <= c
	Where |x| denotes the absolute value of x.

	Return the number of good triplets.


	Example 1:
	Input: arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
	Output: 4
	Explanation: There are 4 good triplets: [(3,0,1), (3,0,1), (3,1,1), (0,1,1)].



*/

package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 3}
	a := 0
	b := 0
	c := 1
	fmt.Println(countGoodTriplets(arr, a, b, c))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var result int
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			for k := j; k < l; k++ {
				if i < j && j < k && k < l && abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					result++
				}
			}
		}
	}

	return result
}
