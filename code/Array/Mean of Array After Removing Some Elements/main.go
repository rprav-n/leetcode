// Author 		: Praveen
// Date   		: 12/01/2023
// Question 	: https://leetcode.com/problems/mean-of-array-after-removing-some-elements/
// Submission 	: https://leetcode.com/problems/mean-of-array-after-removing-some-elements/submissions/876898185/

/*
	Question:
	--------------------------------
	Given an integer array arr, return the mean of the remaining integers after removing the smallest 5% and the largest 5% of the elements.

	Answers within 10-5 of the actual answer will be considered accepted.


	Example 1:
	Input: arr = [1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,3]
	Output: 2.00000
	Explanation: After erasing the minimum and the maximum values of this array, all elements are equal to 2, so the mean is 2.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0}
	// arr = []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}
	fmt.Println(trimMean(arr))
}

func trimMean(arr []int) float64 {

	sort.Ints(arr)
	trim := int(float64(len(arr)) * 0.05)

	temp := arr[trim : len(arr)-trim]

	var sum float64
	for _, i := range temp {
		sum += float64(i)
	}

	return sum / float64(len(temp))
}
