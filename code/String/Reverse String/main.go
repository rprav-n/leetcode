// Author 		: Praveen
// Date   		: 30/04/2023
// Question 	: https://leetcode.com/problems/reverse-string/
// Submission 	: https://leetcode.com/problems/reverse-string/submissions/942029424/

package main

func main() {
	s := []byte("hello")
	reverseString(s)
}

func reverseString(s []byte) {

	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
