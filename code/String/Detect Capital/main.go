// Author 		: Praveen
// Date   		: 23/06/2023
// Question 	: https://leetcode.com/problems/detect-capital/
// Submission 	: https://leetcode.com/problems/detect-capital/submissions/978019371/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	word := "leetcodA"
	fmt.Println(detectCapitalUse(word))
}

func detectCapitalUse(word string) bool {

	allCaps := true
	allNotCaps := true
	onlyFirstLetterIsCap := true

	for i, c := range word {

		isUpper := unicode.IsUpper(c)

		if i == 0 && !isUpper {
			onlyFirstLetterIsCap = false
			allCaps = false
		}

		if i != 0 && isUpper {
			allNotCaps = false
			onlyFirstLetterIsCap = false
		}

		if !isUpper {
			allCaps = false
		}

	}

	if allCaps || allNotCaps || onlyFirstLetterIsCap {
		return true
	}

	return false
}
