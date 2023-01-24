// Author 		: Praveen
// Date   		: 24/01/2023
// Question 	: https://leetcode.com/problems/flipping-an-image/
// Submission 	: https://leetcode.com/problems/flipping-an-image/submissions/884396161/

package main

import "fmt"

func main() {
	image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(image))
}

func flipAndInvertImage(image [][]int) [][]int {
	l := len(image)
	l1 := len(image[0])

	ans := make([][]int, l)

	for i := range ans {
		ans[i] = make([]int, l1)
	}

	for j, n := range image {
		for i := l1 - 1; i >= 0; i-- {
			val := n[i]
			if val == 1 {
				val = 0
			} else {
				val = 1
			}
			ans[j][l1-i-1] = val
		}
	}

	return ans
}
