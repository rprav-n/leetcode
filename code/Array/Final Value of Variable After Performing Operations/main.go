// Author 		: Praveen
// Date   		: 27/11/2022
// Question 	: https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
// Submission 	: https://leetcode.com/problems/final-value-of-variable-after-performing-operations/submissions/850738633/
// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/submissions/850739849/
/*
	Question:
	--------------------------------

	There is a programming language with only four operations and one variable X:

	++X and X++ increments the value of the variable X by 1.
	--X and X-- decrements the value of the variable X by 1.
	Initially, the value of X is 0.

	Given an array of strings operations containing a list of operations, return the final value of X after performing all the operations.

	Example 1:
	Input: operations = ["--X","X++","X++"]
	Output: 1
	Explanation: The operations are performed as follows:
	Initially, X = 0.
	--X: X is decremented by 1, X =  0 - 1 = -1.
	X++: X is incremented by 1, X = -1 + 1 =  0.
	X++: X is incremented by 1, X =  0 + 1 =  1.

	Example 2:
	Input: operations = ["++X","++X","X++"]
	Output: 3
	Explanation: The operations are performed as follows:
	Initially, X = 0.
	++X: X is incremented by 1, X = 0 + 1 = 1.
	++X: X is incremented by 1, X = 1 + 1 = 2.
	X++: X is incremented by 1, X = 2 + 1 = 3.


	Notes:
	--------------------------------

	Simple solution:
	Create a variable of type int (By default int has a zero value ie 0)
	Loop through the operations array string and check if the string contains "+",
	if YES -> Add +1 to result
	else -> Add -1 to result

	Better Solution: Without using strings package Contains method, check if first and last character is "+",
	if YES -> Add +1 to result
	else -> Add -1 to result

*/

package main

import (
	"fmt"
)

func main() {
	operations := []string{"--X", "X++", "X++"}
	fmt.Println(finalValueAfterOperations(operations))
}

func finalValueAfterOperations(operations []string) int {
	var result int
	l := 2
	for _, operation := range operations {
		if operation[0] == byte('+') || operation[l] == byte('+') {
			result++
		} else {
			result--
		}
	}

	return result
}
