// Author 		: Praveen
// Date   		: 12/04/2023
// Question 	: https://leetcode.com/problems/range-sum-query-immutable/
// Submission 	: https://leetcode.com/problems/range-sum-query-immutable/submissions/932623897/

package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	var ans int

	ans = sumArr(this.nums[left : right+1])

	return ans
}

func sumArr(arr []int) int {
	sum := 0

	for _, n := range arr {
		sum += n
	}

	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
