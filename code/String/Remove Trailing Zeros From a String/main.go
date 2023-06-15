// Author 		: Praveen
// Date   		: 15/06/2023
// Question 	: https://leetcode.com/problems/remove-trailing-zeros-from-a-string/
// Submission 	: https://leetcode.com/problems/remove-trailing-zeros-from-a-string/submissions/972081873/

package main

import "fmt"

func main() {
	num := "123"
	fmt.Println(removeTrailingZeros(num))
}

func removeTrailingZeros(num string) string {

	l := len(num)

	for i := l - 1; i >= 0; i-- {
		b := num[i]
		if b != byte('0') {
			num = num[0 : i+1]
			break
		}
	}

	return num
}
