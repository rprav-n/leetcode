// Author 		: Praveen
// Date   		: 21/04/2023
// Question 	: https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet/
// Submission 	: https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet/submissions/937523155/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "A1:F1"
	fmt.Println(cellsInRange(s))
}

func cellsInRange(s string) []string {
	var arr []string

	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	sn, _ := strconv.Atoi(string(s[1]))
	en, _ := strconv.Atoi(string(s[4]))

	startIndex := strings.Index(alphabets, string(s[0]))
	endIndex := strings.Index(alphabets, string(s[3]))

	for j := startIndex; j <= endIndex; j++ {
		letter := string(alphabets[j])
		for i := sn; i <= en; i++ {
			arr = append(arr, fmt.Sprintf("%s%d", letter, i))
		}
	}

	return arr
}
