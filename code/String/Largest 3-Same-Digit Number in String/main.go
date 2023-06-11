// Author 		: Praveen
// Date   		: 11/06/2023
// Question 	: https://leetcode.com/problems/largest-3-same-digit-number-in-string/
// Submission 	: https://leetcode.com/problems/largest-3-same-digit-number-in-string/submissions/968997629/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "42352338"

	fmt.Println(largestGoodInteger(num))
}

func largestGoodInteger(num string) string {
	var ans int

	l := len(num)

	m := make(map[int]struct{})

	for i, c := range num {
		if i < l-2 {
			x := string(c)
			y := string(num[i+1])
			z := string(num[i+2])
			if x == y && y == z {
				ns := x + y + z
				n, _ := strconv.Atoi(ns)
				m[n] = struct{}{}
			}
		}
	}

	if len(m) == 0 {
		return ""
	}

	for k := range m {
		if k > ans {
			ans = k
		}
	}

	if ans == 0 {
		return "000"
	}

	return fmt.Sprintf("%d", ans)
}
