// Author 		: Praveen
// Date   		: 14/12/2022
// Question 	: https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
// Submission 	:

/*
	Question:
	--------------------------------

	Given an array of positive integers arr, return the sum of all possible odd-length subarrays of arr.

	A subarray is a contiguous subsequence of the array.


	Example 1:
	Input: arr = [1,4,2,5,3]
	Output: 58
	Explanation: The odd-length subarrays of arr and their sums are:
	[1] = 1
	[4] = 4
	[2] = 2
	[5] = 5
	[3] = 3
	[1,4,2] = 7
	[4,2,5] = 11
	[2,5,3] = 10
	[1,4,2,5,3] = 15
	If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58

	Example 2:
	Input: arr = [1,2]
	Output: 3
	Explanation: There are only 2 subarrays of odd length, [1] and [2]. Their sum is 3.

	1 2 3 4 5 6

*/

package main

import "fmt"

func main() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays(arr))
}

func sumOddLengthSubarrays(arr []int) int {
	var sum int
	var largestOddLenOfArr int
	l := len(arr)

	if l%2 == 0 {
		largestOddLenOfArr = l - 1
	} else {
		largestOddLenOfArr = l
	}

	for _, n := range arr {
		sum += n
	}

	if l%2 != 0 {
		sum += sum
	}

	for largestOddLenOfArr > 0 {

		for i := 0; i < l; i++ {
			var s, nl int
			if i < largestOddLenOfArr {
				s = getSumOf(arr[i:largestOddLenOfArr])
				nl = len(arr[i:largestOddLenOfArr])
			}
			fmt.Println(i, largestOddLenOfArr, arr[i:largestOddLenOfArr])
			if nl%2 != 0 && s%2 != 0 {
				sum += s
			}
		}

		largestOddLenOfArr -= 2
	}

	// fmt.Println(sum)

	return sum
}

func getSumOf(arr []int) int {
	var s int
	for _, n := range arr {
		s += n
	}
	return s
}
