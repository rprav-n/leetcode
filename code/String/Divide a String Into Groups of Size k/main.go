// Author 		: Praveen
// Date   		: 16/05/2023
// Question 	: https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/
// Submission 	: https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/submissions/951572051/

package main

import "fmt"

func main() {
	s := "ctoyjrwtngqwt"
	k := 8
	fill := byte('x')
	fmt.Println(divideString(s, k, fill))
}

func divideString(s string, k int, fill byte) []string {
	var ans []string
	count := 0
	str := ""
	for _, c := range s {
		str += string(c)
		count++
		if count == k {
			ans = append(ans, str)
			count = 0
			str = ""
		}
	}
	if count != 0 {
		for i := count; i < k; i++ {
			str += string(fill)
		}
		ans = append(ans, str)
	}

	return ans
}
