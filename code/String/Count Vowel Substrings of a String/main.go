// Author 		: Praveen
// Date   		: 15/05/2023
// Question 	: https://leetcode.com/problems/count-vowel-substrings-of-a-string/
// Submission 	: https://leetcode.com/problems/count-vowel-substrings-of-a-string/submissions/951567093/

package main

import "fmt"

func main() {
	word := "cuaieuouac"
	fmt.Println(countVowelSubstrings(word))
}

func countVowelSubstrings(word string) int {
	counter := 0
loop:
	for m := 0; m < len(word); m++ {
		a, e, i, o, u := 0, 0, 0, 0, 0
		for n := m; n < len(word); n++ {
			switch word[n] {
			case 'a':
				a++
			case 'e':
				e++
			case 'i':
				i++
			case 'o':
				o++
			case 'u':
				u++
			default:
				continue loop
			}
			if a > 0 && e > 0 && i > 0 && o > 0 && u > 0 {
				counter++
			}
		}
	}
	return counter
}
