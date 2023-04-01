// Author 		: Praveen
// Date   		: 01/04/2023
// Question 	: https://leetcode.com/problems/teemo-attacking/
// Submission 	: https://leetcode.com/problems/teemo-attacking/submissions/926076492/

package main

import (
	"fmt"
)

func main() {
	timeSeries := []int{1, 2}
	duration := 2
	fmt.Println(findPoisonedDuration(timeSeries, duration))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0

	for i := 0; i < len(timeSeries)-1; i++ {
		total += min(timeSeries[i+1]-timeSeries[i], duration)
	}

	return total + duration
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
