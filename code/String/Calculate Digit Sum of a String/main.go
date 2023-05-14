// Author 		: Praveen
// Date   		: 14/05/2023
// Question 	: https://leetcode.com/problems/calculate-digit-sum-of-a-string/
// Submission 	: https://leetcode.com/problems/calculate-digit-sum-of-a-string/submissions/950091843/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "1234"
	k := 2
	fmt.Println(digitSum(s, k))
}

func digitSum(s string, k int) string {

	for len(s) > k {
		numsStr := []string{}
		count := 0
		ns := ""
		for _, c := range s {
			ns += string(c)
			count++
			if count == k {
				numsStr = append(numsStr, ns)
				count = 0
				ns = ""
			}
		}
		if ns != "" {
			numsStr = append(numsStr, ns)
		}

		fmt.Println(numsStr)
		s = ""
		for _, numStr := range numsStr {
			sum := 0
			for _, c := range numStr {
				n, _ := strconv.Atoi(string(c))
				sum += n
			}
			s += fmt.Sprintf("%d", sum)
		}
		fmt.Println(s)
	}

	return s
}
