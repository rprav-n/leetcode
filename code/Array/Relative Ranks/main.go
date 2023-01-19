// Author 		: Praveen
// Date   		: 19/01/2023
// Question 	: https://leetcode.com/problems/relative-ranks/
// Submission 	: https://leetcode.com/problems/relative-ranks/submissions/881313966/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	score := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))
}

func findRelativeRanks(score []int) []string {
	var ans []string
	l := len(score)
	cp := make([]int, l)
	m := make(map[int]int)

	copy(cp, score)
	sort.Ints(score)

	r := l
	for _, s := range score {
		m[s] = r
		r--
	}

	for _, n := range cp {
		rank := m[n]
		if rank == 1 {
			ans = append(ans, "Gold Medal")
		} else if rank == 2 {
			ans = append(ans, "Silver Medal")
		} else if rank == 3 {
			ans = append(ans, "Bronze Medal")
		} else {
			ans = append(ans, strconv.Itoa(rank))
		}
	}

	return ans
}
