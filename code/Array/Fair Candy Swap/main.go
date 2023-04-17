// Author 		: Praveen
// Date   		: 17/04/2023
// Question 	: https://leetcode.com/problems/fair-candy-swap/
// Submission 	: https://leetcode.com/problems/fair-candy-swap/submissions/935029501/

package main

import "fmt"

func main() {
	aliceSizes := []int{2}
	bobSizes := []int{1, 3}
	fmt.Println(fairCandySwap(aliceSizes, bobSizes))
}

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	var res []int

	sumOfAlice := sumOff(aliceSizes)
	sumOfBob := sumOff(bobSizes)

	for i := 0; i < len(aliceSizes); i++ {
		for j := 0; j < len(bobSizes); j++ {

			newSumAlice := sumOfAlice - aliceSizes[i]
			newSumBob := sumOfBob - bobSizes[j]

			newSumAlice += bobSizes[j]
			newSumBob += aliceSizes[i]

			if newSumAlice == newSumBob {
				res = append(res, aliceSizes[i], bobSizes[j])
				return res
			}

		}
	}

	return res
}

func sumOff(arr []int) int {
	s := 0
	for _, n := range arr {
		s += n
	}

	return s
}
