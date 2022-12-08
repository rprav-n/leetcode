// Author 		: Praveen
// Date   		: 08/12/2022
// Question 	: https://leetcode.com/problems/decompress-run-length-encoded-list/
// Submission 	: https://leetcode.com/problems/decompress-run-length-encoded-list/submissions/856710771/

/*
	Question:
	--------------------------------
	We are given a list nums of integers representing a list compressed with run-length encoding.

	Consider each adjacent pair of elements [freq, val] = [nums[2*i], nums[2*i+1]] (with i >= 0).  For each such pair, there are freq elements with value val concatenated in a sublist. Concatenate all the sublists from left to right to generate the decompressed list.

	Return the decompressed list.


	Example 1:
	Input: nums = [1,2,3,4]
	Output: [2,4,4,4]
	Explanation: The first pair [1,2] means we have freq = 1 and val = 2 so we generate the array [2].
	The second pair [3,4] means we have freq = 3 and val = 4 so we generate [4,4,4].
	At the end the concatenation [2] + [4,4,4] is [2,4,4,4].

	Example 2:
	Input: nums = [1,1,2,3]
	Output: [1,3,3]


	Notes:
	--------------------------------


*/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3}
	fmt.Println(decompressRLElist(nums))
}

func decompressRLElist(nums []int) []int {
	var result []int

	for i, n := range nums {
		if i%2 == 0 {
			j := n
			for j > 0 {
				result = append(result, nums[i+1])
				j--
			}
		}
	}

	return result
}
