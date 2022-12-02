// Author 		: Praveen
// Date   		: 02/12/2022
// Question 	: https://leetcode.com/problems/richest-customer-wealth
// Submission 	: https://leetcode.com/problems/richest-customer-wealth/submissions/853519986/

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
