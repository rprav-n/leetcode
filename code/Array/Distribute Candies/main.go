// Author 		: Praveen
// Date   		: 11/01/2023
// Question 	: https://leetcode.com/problems/distribute-candies/
// Submission 	:

/*
	Question:
	--------------------------------
	Alice has n candies, where the ith candy is of type candyType[i]. Alice noticed that she started to gain weight, so she visited a doctor.

	The doctor advised Alice to only eat n / 2 of the candies she has (n is always even). Alice likes her candies very much, and she wants to eat the maximum number of different types of candies while still following the doctor's advice.

	Given the integer array candyType of length n, return the maximum number of different types of candies she can eat if she only eats n / 2 of them.


	Example 1:
	Input: candyType = [1,1,2,2,3,3]
	Output: 3
	Explanation: Alice can only eat 6 / 2 = 3 candies. Since there are only 3 types, she can eat one of each type.

*/

package main

import "fmt"

func main() {
	candyType := []int{6, 6, 6, 6}
	fmt.Println(distributeCandies(candyType))
}

func distributeCandies(candyType []int) int {
	var ans int

	m := make(map[int]struct{})

	for _, n := range candyType {
		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
		}
	}

	l := len(candyType)
	canHave := l / 2

	if len(m) >= canHave {
		ans = canHave
	} else {
		ans = len(m)
	}

	return ans
}
