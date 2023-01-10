// Author 		: Praveen
// Date   		: 10/01/2023
// Question 	: https://leetcode.com/problems/number-of-lines-to-write-string/
// Submission 	: https://leetcode.com/problems/number-of-lines-to-write-string/submissions/875637837/

/*
	Question:
	--------------------------------
	You are given a string s of lowercase English letters and an array widths denoting how many pixels wide each lowercase English letter is. Specifically, widths[0] is the width of 'a', widths[1] is the width of 'b', and so on.

	You are trying to write s across several lines, where each line is no longer than 100 pixels. Starting at the beginning of s, write as many letters on the first line such that the total width does not exceed 100 pixels. Then, from where you stopped in s, continue writing as many letters as you can on the second line. Continue this process until you have written all of s.

	Return an array result of length 2 where:

	result[0] is the total number of lines.
	result[1] is the width of the last line in pixels.



	Example 1:
	Input: widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10], s = "abcdefghijklmnopqrstuvwxyz"
	Output: [3,60]
	Explanation: You can write s as follows:
	abcdefghij  // 100 pixels wide
	klmnopqrst  // 100 pixels wide
	uvwxyz      // 60 pixels wide
	There are a total of 3 lines, and the last line is 60 pixels wide.


*/

package main

import "fmt"

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"

	widths = []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s = "bbbcccdddaaa"
	widths = []int{3, 4, 10, 4, 8, 7, 3, 3, 4, 9, 8, 2, 9, 6, 2, 8, 4, 9, 9, 10, 2, 4, 9, 10, 8, 2}
	s = "mqblbtpvicqhbrejb"

	fmt.Println(numberOfLines(widths, s))
}

func numberOfLines(widths []int, s string) []int {

	m := make(map[rune]int)
	alphabets := "abcdefghijklmnopqrstuvwxyz"

	for i, r := range alphabets {
		m[r] = widths[i]
	}

	lines := 0
	count := 0

	for _, ch := range s {
		if count+m[ch] > 100 {
			lines++
			count = m[ch]
		} else {
			count += m[ch]
		}
	}

	return []int{lines + 1, count}
}
