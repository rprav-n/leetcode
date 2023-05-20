// Author 		: Praveen
// Date   		: 20/05/2023
// Question 	: https://leetcode.com/problems/reformat-phone-number/
// Submission 	: https://leetcode.com/problems/reformat-phone-number/submissions/954050358/

package main

import (
	"fmt"
	"strings"
)

func main() {
	number := "123 4-5678"
	fmt.Println(reformatNumber(number))
}

func reformatNumber(number string) string {
	var ans string

	num := strings.ReplaceAll(number, "-", "")
	num = strings.ReplaceAll(num, " ", "")

	for len(num) != 0 {
		l := len(num)
		if l > 4 {
			ans += num[0:3]
			ans += "-"
			num = num[3:]
		} else if l == 4 {
			ans += num[0:2] + "-" + num[2:4]
			break
		} else {
			ans += num[:]
			break
		}
	}

	return ans
}
