// Author 		: Praveen
// Date   		: 24/01/2023
// Question 	: https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
// Submission 	: https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/submissions/884389002/

package main

import (
	"fmt"
	"sort"
)

func main() {
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}
	fmt.Println(minMovesToSeat(seats, students))
}

func minMovesToSeat(seats []int, students []int) int {
	var ans int
	l := len(seats)

	sort.Ints(seats)
	sort.Ints(students)

	for i := 0; i < l; i++ {
		a := seats[i]
		b := students[i]
		ans += Abs(a - b)
	}

	return ans
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
	1 2 3 6
	1 4 5 9
	0 2 2 3

	1 3 5
	2 4 7
	1 1 2
*/
