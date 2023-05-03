// Author 		: Praveen
// Date   		: 03/05/2023
// Question 	: https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
// Submission 	: https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/submissions/943865114/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	firstWord := "acb"
	secondWord := "cba"
	targetWord := "cdb"

	fmt.Println(isSumEqual(firstWord, secondWord, targetWord))
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {

	alphabets := "abcdefghijklmnopqrstuvwxyz"

	am := make(map[rune]int)

	for i, c := range alphabets {
		am[c] = i
	}

	var fs, ss, ts string

	for _, c := range firstWord {
		fs += fmt.Sprintf("%d", am[c])
	}
	for _, c := range secondWord {
		ss += fmt.Sprintf("%d", am[c])
	}
	for _, c := range targetWord {
		ts += fmt.Sprintf("%d", am[c])
	}

	fsi, _ := strconv.Atoi(fs)
	ssi, _ := strconv.Atoi(ss)
	tsi, _ := strconv.Atoi(ts)

	if fsi+ssi == tsi {
		return true
	}

	return false
}
