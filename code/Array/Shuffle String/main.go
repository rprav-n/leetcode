// Author 		: Praveen
// Date   		: 10/12/2022
// Question 	: https://leetcode.com/problems/shuffle-string/
// Submission 	: https://leetcode.com/problems/shuffle-string/submissions/857686844/

/*
	Question:
	--------------------------------
	You are given a string s and an integer array indices of the same length. The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

	Return the shuffled string.


	Example 1:
	Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
	Output: "leetcode"
	Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.

	Example 2:
	Input: s = "abc", indices = [0,1,2]
	Output: "abc"
	Explanation: After shuffling, each character remains in its position.


	Notes:
	--------------------------------


*/

package main

import (
	"fmt"
)

func main() {
	s := "abc"
	indices := []int{0, 1, 2}
	fmt.Println(restoreString(s, indices))

}

func restoreString(s string, indices []int) string {
	var r string
	l := len(indices)
	m := make(map[int]byte)

	for i, c := range s {
		m[indices[i]] = byte(c)
	}

	for i := 0; i < l; i++ {
		r += string(m[i])
	}

	return r
}
