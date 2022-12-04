// Author 		: Praveen
// Date   		: 02/12/2022
// Question 	: https://leetcode.com/problems/richest-customer-wealth
// Submission 	: https://leetcode.com/problems/richest-customer-wealth/submissions/853519986/

/*
	Question:
	--------------------------------

	You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.

	A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.

	Example 1:
	Input: accounts = [[1,2,3],[3,2,1]]
	Output: 6
	Explanation:
	1st customer has wealth = 1 + 2 + 3 = 6
	2nd customer has wealth = 3 + 2 + 1 = 6
	Both customers are considered the richest with a wealth of 6 each, so return 6.

	Example 2:
	Input: accounts = [[1,5],[7,3],[3,5]]
	Output: 10
	Explanation:
	1st customer has wealth = 6
	2nd customer has wealth = 10
	3rd customer has wealth = 8
	The 2nd customer is the richest with a wealth of 10.


	Notes:
	--------------------------------
	Have a variable called "max" of type int which will be our end result
	Loop (i) throught the main arr ie accounts (Nested array)
		Get the sum of the nested array ie accounts[i]
		if sum > max:
		then: assign max = sum
	return the max result

*/

package main

import "fmt"

func main() {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(accounts))
}

func maximumWealth(accounts [][]int) int {
	var max int

	for _, account := range accounts {
		sum := getSumOfArr(account)
		if sum > max {
			max = sum
		}
	}

	return max
}

func getSumOfArr(arr []int) int {
	var sum int

	for _, val := range arr {
		sum += val
	}

	return sum
}
