// Author 		: Praveen
// Date   		: 26/01/2023
// Question 	: https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
// Submission 	: https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/submissions/885561767/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	fmt.Println(sortByBits(arr))

}

func sortByBits(arr []int) []int {

	sort.Slice(arr, func(i, j int) bool {
		a := getNumberOf1Bits(arr[i])
		b := getNumberOf1Bits(arr[j])
		if a != b {
			return a < b
		}
		return arr[i] < arr[j]
	})

	return arr
}

func getNumberOf1Bits(n int) int {
	var count int
	i := int64(n)
	s := strconv.FormatInt(i, 2)

	for _, c := range s {
		if c == '1' {
			count++
		}
	}

	return count
}
