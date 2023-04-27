// Author 		: Praveen
// Date   		: 27/04/2023
// Question 	: https://leetcode.com/problems/reverse-prefix-of-word/
// Submission 	: https://leetcode.com/problems/reverse-prefix-of-word/submissions/940649672/

package main

import "fmt"

func main() {
	word := "xyxzxe"
	ch := byte(rune('z'))

	fmt.Println(reversePrefix(word, ch))
}

func reversePrefix(word string, ch byte) string {

	var ans string

	for i, c := range word {

		if c == rune(ch) {
			for j := i; j >= 0; j-- {
				ans += string(word[j])
			}
			for k := i + 1; k < len(word); k++ {
				ans += string(word[k])
			}
			break
		}
	}
	if ans == "" {
		ans = word
	}

	return ans

}
