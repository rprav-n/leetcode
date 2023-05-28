// Author 		: Praveen
// Date   		: 28/05/2023
// Question 	: https://leetcode.com/problems/capitalize-the-title/
// Submission 	: https://leetcode.com/problems/capitalize-the-title/submissions/959107128/

package main

import (
	"fmt"
	"strings"
)

func main() {
	title := "i lOve leetcode"
	fmt.Println(capitalizeTitle(title))
}

func capitalizeTitle(title string) string {
	var ans string

	words := strings.Split(title, " ")

	l := len(words)

	for i, word := range words {

		if len(word) <= 2 {
			ans += strings.ToLower(word)
		} else {
			word = strings.ToLower(word)
			ans += strings.Title(word)
		}
		if i < l-1 {
			ans += " "
		}
	}

	return ans
}
