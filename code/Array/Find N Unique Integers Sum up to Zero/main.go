// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
// Submission 	: https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/submissions/866106234/

/*
	Question:
	--------------------------------
	Given an integer n, return any array containing n unique integers such that they add up to 0.


	Example 1:
	Input: n = 5
	Output: [-7,-1,1,3,4]
	Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].


*/

package main

import "fmt"

func main() {
	n := 1
	fmt.Println(sumZero(n))
}

func sumZero(n int) []int {
	var arr []int
	var s int
	for i := 1; i < n; i++ {
		s += i
		arr = append(arr, i)
	}

	arr = append(arr, -s)

	return arr
}
