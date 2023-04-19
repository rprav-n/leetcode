// Author 		: Praveen
// Date   		: 19/04/2023
// Question 	: https://leetcode.com/problems/jewels-and-stones/
// Submission 	: https://leetcode.com/problems/jewels-and-stones/submissions/936468259/

package main

import "fmt"

func main() {
	jeweles := "z"
	stones := "ZZ"
	fmt.Println(numJewelsInStones(jeweles, stones))

}

func numJewelsInStones(jewels string, stones string) int {
	var count int

	m := make(map[rune]struct{})

	for _, r := range jewels {
		m[r] = struct{}{}
	}

	for _, r := range stones {
		if _, ok := m[r]; ok {
			count++
		}
	}

	return count
}
