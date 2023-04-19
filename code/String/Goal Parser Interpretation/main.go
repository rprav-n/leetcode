// Author 		: Praveen
// Date   		: 19/04/2023
// Question 	: https://leetcode.com/problems/goal-parser-interpretation/
// Submission 	: https://leetcode.com/problems/goal-parser-interpretation/submissions/936474666/

package main

import "fmt"

func main() {
	command := "G()()()()(al)"
	fmt.Println(interpret(command))
}

func interpret(command string) string {
	var ans string

	l := len(command)

	for i := 0; i < l; i++ {
		if command[i] == 'G' {
			ans += "G"
		} else if (i+1) < l && command[i] == '(' && command[i+1] == ')' {
			i = i + 1
			ans += "o"
		} else if (i+1) < l && command[i] == '(' && command[i+1] == 'a' {
			i = i + 2
			ans += "al"
		}
	}

	return ans
}
