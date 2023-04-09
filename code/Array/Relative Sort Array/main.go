// Author 		: Praveen
// Date   		: 09/04/2023
// Question 	: https://leetcode.com/problems/relative-sort-array/
// Submission 	: https://leetcode.com/problems/relative-sort-array/submissions/930624762/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{28, 6, 22, 8, 44, 17}
	arr2 := []int{22, 28, 8, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var ans []int
	arr2Map := make(map[int]struct{})

	for _, n := range arr2 {
		arr2Map[n] = struct{}{}
	}

	for _, n := range arr2 {
		for _, m := range arr1 {
			if n == m {
				ans = append(ans, n)
			}
		}
	}

	// fmt.Println(arr1, arr2Map)
	newArr := []int{}
	for _, n := range arr1 {
		if _, ok := arr2Map[n]; !ok {
			newArr = append(newArr, n)
		}
	}

	sort.Ints(newArr)
	ans = append(ans, newArr...)

	return ans
}
