// Author 		: Praveen
// Date   		: 28/04/2023
// Question 	: https://leetcode.com/problems/determine-color-of-a-chessboard-square/
// Submission 	: https://leetcode.com/problems/determine-color-of-a-chessboard-square/submissions/941221246/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	coordinates := "a1"
	fmt.Println(squareIsWhite(coordinates))
}

func squareIsWhite(coordinates string) bool {

	alphabets := "abcdefgh"

	letterPos := 0
	numberPos, _ := strconv.Atoi(string(coordinates[1]))

	for i, c := range alphabets {
		if c == rune(coordinates[0]) {
			letterPos = i + 1
			break
		}
	}

	if (letterPos%2 == 0 && numberPos%2 != 0) || (letterPos%2 != 0 && numberPos%2 == 0) {
		return true
	}

	return false
}
