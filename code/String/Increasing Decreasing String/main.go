// Author 		: Praveen
// Date   		: 30/04/2023
// Question 	: https://leetcode.com/problems/increasing-decreasing-string/
// Submission 	: https://leetcode.com/problems/increasing-decreasing-string/submissions/942027135/

package main

import "fmt"

func main() {
	s := "rat"
	fmt.Println(sortString(s))
}

func sortString(s string) string {
	var ans string

	alphabets := "abcdefghijklmnopqrstuvwxyz"

	f := make(map[rune]int)

	for _, c := range s {
		if _, ok := f[c]; ok {
			f[c] += 1
		} else {
			f[c] = 1
		}
	}

	for !checkIfFrequencyIsZero(f) {
		for _, c := range alphabets {
			if val, ok := f[c]; ok && val > 0 {
				ans += string(c)
				f[c] -= 1
			}
		}
		for i := 25; i >= 0; i-- {
			c := rune(alphabets[i])
			if val, ok := f[c]; ok && val > 0 {
				ans += string(c)
				f[c] -= 1
			}
		}
	}

	return ans
}

func checkIfFrequencyIsZero(f map[rune]int) bool {
	for _, v := range f {
		if v > 0 {
			return false
		}
	}

	return true
}
