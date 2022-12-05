// Author 		: Praveen
// Date   		: 05/12/2022
// Question 	: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Submission 	: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/submissions/855137070/

/*
	Question:
	--------------------------------
	There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

	Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

	Note that multiple kids can have the greatest number of candies.


	Example 1:
	Input: candies = [2,3,5,1,3], extraCandies = 3
	Output: [true,true,true,false,true]
	Explanation: If you give all extraCandies to:
	- Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
	- Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
	- Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
	- Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
	- Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.

	Example 2:
	Input: candies = [4,2,1,1,2], extraCandies = 1
	Output: [true,false,false,false,false]
	Explanation: There is only 1 extra candy.
	Kid 1 will always have the greatest number of candies, even if a different kid is given the extra candy.


	Notes:
	--------------------------------
	1. Get the largest number in the array ie candies
	2. Loop through the candies
		2.1 If extracandy + currentLoop candy >= largest number
			then -> Append "True" to results array of bool
			else -> Append "False" to results array of bool
	3. Return the result

*/

package main

import "fmt"

func main() {
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result []bool

	max := getMax(candies)

	for _, candy := range candies {
		if candy+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}

func getMax(candies []int) int {
	var max int

	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	return max
}
