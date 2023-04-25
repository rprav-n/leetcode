// Author 		: Praveen
// Date   		: 25/04/2023
// Question 	: https://leetcode.com/problems/rings-and-rods/
// Submission 	: https://leetcode.com/problems/rings-and-rods/submissions/939645649/

package main

import (
	"fmt"
	"strings"
)

func main() {
	rings := "B9R9G0R4G6R8R2R9G9"
	fmt.Println(countPoints(rings))
}

func countPoints(rings string) int {
	var ans int

	m := make(map[string]string)

	l := len(rings)

	for i := 1; i < l; i += 2 {
		color := string(rings[i-1])
		val := string(rings[i])

		m[val] += color
	}
	fmt.Println(m)

	for _, v := range m {
		if len(v) >= 3 && strings.Contains(v, "R") && strings.Contains(v, "G") && strings.Contains(v, "B") {
			ans++
		}
	}

	return ans
}
