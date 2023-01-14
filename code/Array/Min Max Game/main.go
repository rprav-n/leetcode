// Author 		: Praveen
// Date   		: 14/01/2023
// Question 	: https://leetcode.com/problems/min-max-game/
// Submission 	: https://leetcode.com/problems/min-max-game/submissions/878079511/

/*
	Question:
	--------------------------------
	You are given a 0-indexed integer array nums whose length is a power of 2.

	Apply the following algorithm on nums:

	Let n be the length of nums. If n == 1, end the process. Otherwise, create a new 0-indexed integer array newNums of length n / 2.
	For every even index i where 0 <= i < n / 2, assign the value of newNums[i] as min(nums[2 * i], nums[2 * i + 1]).
	For every odd index i where 0 <= i < n / 2, assign the value of newNums[i] as max(nums[2 * i], nums[2 * i + 1]).
	Replace the array nums with newNums.
	Repeat the entire process starting from step 1.
	Return the last number that remains in nums after applying the algorithm.


	Example 1:
	Input: nums = [1,3,5,2,4,8,2,2]
	Output: 1
	Explanation: The following arrays are the results of applying the algorithm repeatedly.
	First: nums = [1,5,4,2]
	Second: nums = [1,4]
	Third: nums = [1]
	1 is the last remaining number, so we return 1.


*/

package main

import "fmt"

func main() {
	nums := []int{70, 38, 21, 22}
	fmt.Println(minMaxGame(nums))
}

func minMaxGame(nums []int) int {
	var ans int

	ans = calc(nums)

	return ans
}

func calc(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := len(nums)
	c := 0
	arr := []int{}
	for i := 0; i < l; i += 2 {
		n1 := nums[i]
		n2 := nums[i+1]
		if c%2 == 0 {
			arr = append(arr, min(n1, n2))
		} else {
			arr = append(arr, max(n1, n2))
		}
		c++
	}
	fmt.Println(arr)
	return calc(arr)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
