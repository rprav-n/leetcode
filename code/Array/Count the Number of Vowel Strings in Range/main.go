// Author 		: Praveen
// Date   		: 08/04/2023
// Question 	: https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/
// Submission 	: https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/submissions/930066466/

package main

import "fmt"

func main() {
	words := []string{"hey", "aeo", "mu", "ooo", "artro"}
	left := 1
	right := 4
	fmt.Println(vowelStrings(words, left, right))
}

func vowelStrings(words []string, left int, right int) int {
	var res int
	vowels := map[string]struct{}{}
	vowels["a"] = struct{}{}
	vowels["e"] = struct{}{}
	vowels["i"] = struct{}{}
	vowels["o"] = struct{}{}
	vowels["u"] = struct{}{}
	for i, word := range words {
		if i >= left && i <= right {
			l := len(word)
			_, ok1 := vowels[string(word[0])]
			_, ok2 := vowels[string(word[l-1])]
			if ok1 && ok2 {
				res++
			}
		}
	}

	return res
}
