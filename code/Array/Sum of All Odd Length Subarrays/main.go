// Author 		: Praveen
// Date   		: 14/12/2022
// Question 	: https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
// Submission 	: https://leetcode.com/problems/sum-of-all-odd-length-subarrays/submissions/883872749/

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
	arr := []int{10, 11, 12}
	fmt.Println(sumOddLengthSubarrays(arr))
}

func sumOddLengthSubarrays(arr []int) int {
	var mainSum int

	l := len(arr)
	possibleOdds := []int{}
	for i := 0; i <= l; i++ {
		if i%2 != 0 {
			possibleOdds = append(possibleOdds, i)
		}
	}

	for _, n := range possibleOdds {
		sum := 0
		// SLIDING WINDOW
		for i := 0; i < n; i++ {
			sum += arr[i]
		}
		fmt.Println(sum)
		oldSum := sum
		for j := n; j < l; j++ {
			oldSum = oldSum - arr[j-n] + arr[j]
			sum += oldSum
		}
		mainSum += sum

		// Naive solution
		/* for i := 0; i <= l-n; i++ {
			a := arr[i : n+i]
			if len(arr) == 1 {
				sum += a[0]
			} else {
				sum += sumOf(a)
			}
		} */
	}

	return mainSum
}

/* func sumOf(arr []int) int {
	var sum = 0
	for _, n := range arr {
		sum += n
	}

	return sum
} */
