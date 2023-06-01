// Author 		: Praveen
// Date   		: 01/06/2023
// Question 	: https://leetcode.com/problems/consecutive-characters/
// Submission 	: https://leetcode.com/problems/consecutive-characters/submissions/961818901/

package main

import (
	"fmt"
)

func main() {
	s := "yxdxwnjstjqqdgauhnnqcxxulvbulumwfrmgbbplfummzpftoswjlskhzvyidufziaabcygmtiyksmxzaapdwezyzwvqwfffxcqxmgixjendhjmptljcptfupktrmonffzfsndcggihxjfhlgdqnnbvqxizviskqlupuquikfnprylrjrzptvwuvwguhmypqwwgfthjeuqjhqpauliafrjoqyrecxyvuougm"
	fmt.Println(maxPower(s))
}

func maxPower(s string) int {

	l := len(s)
	if l == 1 {
		return 1
	}

	start := s[0]

	count := 1
	odlCount := 1
	for i := 1; i < l; i++ {
		// fmt.Println(string(start), string(s[i]))
		if start == s[i] {
			count++
		} else {
			count = 1
		}
		if count > odlCount {
			odlCount = count
		}
		start = s[i]
	}

	return odlCount
}
