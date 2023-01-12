// Author 		: Praveen
// Date   		: 12/01/2023
// Question 	: https://leetcode.com/problems/last-stone-weight/
// Submission 	: https://leetcode.com/problems/last-stone-weight/submissions/876916417/

/*
	Question:
	--------------------------------
	You are given an array of integers stones where stones[i] is the weight of the ith stone.

	We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:

	If x == y, both stones are destroyed, and
	If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
	At the end of the game, there is at most one stone left.

	Return the weight of the last remaining stone. If there are no stones left, return 0.


	Example 1:
	Input: stones = [2,7,4,1,8,1]
	Output: 1
	Explanation:
	We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
	we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
	we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
	we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	stones := []int{1}
	fmt.Println(lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	var ans int

	sort.Ints(stones)
	fmt.Println(stones)
	l := len(stones)

	for len(stones) > 1 {

		y := stones[l-1]
		x := stones[l-2]
		fmt.Println(x, y)
		if x == y {
			stones = stones[:l-2]
		} else if x != y {
			stones[l-1] = y - x
			stones = remove(stones, l-2)
		}
		sort.Ints(stones)
		l = len(stones)
	}

	if len(stones) == 0 {
		ans = 0
	} else {
		ans = stones[0]
	}
	return ans
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
