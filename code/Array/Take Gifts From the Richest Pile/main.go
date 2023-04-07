// Author 		: Praveen
// Date   		: 07/04/2023
// Question 	: https://leetcode.com/problems/take-gifts-from-the-richest-pile/
// Submission 	: https://leetcode.com/problems/take-gifts-from-the-richest-pile/submissions/929467536/

package main

import (
	"fmt"
	"math"
)

func main() {
	gifts := []int{1, 1, 1, 1}
	k := 4

	fmt.Println(pickGifts(gifts, k))
}

func pickGifts(gifts []int, k int) int64 {
	var ans int64

	s := 0
	for s < k {
		i := getMaxNumIndex(gifts)
		gifts[i] = int(math.Sqrt(float64(gifts[i])))
		s++
	}

	for _, n := range gifts {
		ans += int64(n)
	}

	return ans
}

func getMaxNumIndex(arr []int) int {
	var index int
	s := 0
	for i, n := range arr {
		if n > s {
			s = n
			index = i
		}
	}

	return index
}
