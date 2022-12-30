// Author 		: Praveen
// Date   		: 30/12/2022
// Question 	: https://leetcode.com/problems/two-out-of-three/
// Submission 	: https://leetcode.com/problems/two-out-of-three/submissions/868048606/

/*
	Question:
	--------------------------------
	Given three integer arrays nums1, nums2, and nums3, return a distinct array containing all the values that are present in at least two out of the three arrays. You may return the values in any order.


	Example 1:
	Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
	Output: [3,2]
	Explanation: The values that are present in at least two arrays are:
	- 3, in all three arrays.
	- 2, in nums1 and nums2.


*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 1, 3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{3}

	fmt.Println(twoOutOfThree(nums1, nums2, nums3))
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var result []int
	rm := make(map[int]struct{})

	foo(rm, nums1, nums2, nums3)
	foo(rm, nums2, nums1, nums3)
	foo(rm, nums3, nums2, nums1)
	for k := range rm {
		result = append(result, k)
	}
	return result

}

func foo(rm map[int]struct{}, nums1, nums2, nums3 []int) {
	for _, n := range nums1 {
		for _, v := range nums2 {
			if n == v {
				rm[n] = struct{}{}
			}
		}
		for _, v := range nums3 {
			if n == v {
				rm[n] = struct{}{}
			}
		}
	}
}
