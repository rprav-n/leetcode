// Author 		: Praveen
// Date   		: 22/03/2023
// Question 	: https://leetcode.com/problems/decode-the-message/
// Submission 	: https://leetcode.com/problems/decode-the-message/submissions/938003743/

package main

import (
	"fmt"
	"strings"
)

func main() {
	key := "eljuxhpwnyrdgtqkviszcfmabo"
	message := "zwx hnfx lqantp mnoeius ycgk vcnjrdb"
	fmt.Println(decodeMessage(key, message))
}

func decodeMessage(key string, message string) string {
	var ans string

	alphabtes := "abcdefghijklmnopqrstuvwxyz"
	newKey := strings.ReplaceAll(key, " ", "")

	m := make(map[rune]byte)

	i := 0
	for _, c := range newKey {
		if _, ok := m[c]; !ok {
			m[c] = alphabtes[i]
			i++
		}
	}

	for _, c := range message {
		if c == rune(' ') {
			ans += " "
		} else {
			ans += string(m[c])
		}

	}

	return ans
}
