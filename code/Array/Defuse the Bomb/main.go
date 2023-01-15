// Author 		: Praveen
// Date   		: 15/01/2023
// Question 	: https://leetcode.com/problems/defuse-the-bomb/
// Submission 	: https://leetcode.com/problems/defuse-the-bomb/submissions/878739594/

/*
	Question:
	--------------------------------
	You have a bomb to defuse, and your time is running out! Your informer will provide you with a circular array code of length of n and a key k.

	To decrypt the code, you must replace every number. All the numbers are replaced simultaneously.

	If k > 0, replace the ith number with the sum of the next k numbers.
	If k < 0, replace the ith number with the sum of the previous k numbers.
	If k == 0, replace the ith number with 0.
	As code is circular, the next element of code[n-1] is code[0], and the previous element of code[0] is code[n-1].

	Given the circular array code and an integer key k, return the decrypted code to defuse the bomb!


	Example 1:
	Input: code = [5,7,1,4], k = 3
	Output: [12,10,16,13]
	Explanation: Each number is replaced by the sum of the next 3 numbers. The decrypted code is [7+1+4, 1+4+5, 4+5+7, 5+7+1]. Notice that the numbers wrap around.


*/

package main

import (
	"fmt"
)

func main() {
	code := []int{5, 7, 1, 4}
	//			  0, 1, 2, 3
	k := 3
	fmt.Println(decrypt(code, k))

}

func decrypt(code []int, k int) []int {
	var ans []int
	l := len(code)
	ck := k
	if k < 0 {
		code = append(reverseInts(code[1:]), code[0])
		ck = -k
	}

	for i := 0; i < l; i++ {
		j := 0
		indices := []int{}

		if k != 0 {
			for j < ck {
				v := j + i + 1
				if v > l-1 {
					indices = append(indices, v-l)
				} else {
					indices = append(indices, j+i+1)
				}

				j++
			}

			ans = append(ans, sumOf(indices, code))

		} else {
			ans = append(ans, 0)
		}
	}

	if k < 0 {
		ans = append(reverseInts(ans[1:]), ans[0])
	}

	return ans
}

func sumOf(indices, code []int) int {
	var sum int

	for _, in := range indices {
		sum += code[in]
	}

	return sum
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
