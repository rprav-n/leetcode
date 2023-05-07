// Author 		: Praveen
// Date   		: 07/05/2023
// Question 	: https://leetcode.com/problems/fizz-buzz/
// Submission 	: https://leetcode.com/problems/fizz-buzz/submissions/946097221/

package main

import "fmt"

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	var ans []string

	for i := 1; i <= n; i++ {
		div3 := i % 3
		div5 := i % 5
		if div3 == 0 && div5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if div3 == 0 {
			ans = append(ans, "Fizz")
		} else if div5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, fmt.Sprintf("%d", i))
		}
	}

	return ans
}
