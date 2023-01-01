// Author 		: Praveen
// Date   		: 01/01/2023
// Question 	: https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/
// Submission 	: https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/submissions/869122115/

/*
	Question:
	--------------------------------
	The value of an alphanumeric string can be defined as:

	The numeric representation of the string in base 10, if it comprises of digits only.
	The length of the string, otherwise.
	Given an array strs of alphanumeric strings, return the maximum value of any string in strs.


	Example 1:
	Input: strs = ["alic3","bob","3","4","00000"]
	Output: 5
	Explanation:
	- "alic3" consists of both letters and digits, so its value is its length, i.e. 5.
	- "bob" consists only of letters, so its value is also its length, i.e. 3.
	- "3" consists only of digits, so its value is its numeric equivalent, i.e. 3.
	- "4" also consists only of digits, so its value is 4.
	- "00000" consists only of digits, so its value is 0.
	Hence, the maximum value is 5, of "alic3".


*/

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	strs := []string{"1", "01", "001", "0001"}
	fmt.Println(maximumValue(strs))
}

func maximumValue(strs []string) int {
	var ans int

	for _, s := range strs {
		val := 0
		if regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(s) {
			val = len(s)
		} else if regexp.MustCompile(`^[0-9]+$`).MatchString(s) {
			val, _ = strconv.Atoi(s)
		} else {
			val = len(s)
		}

		if val > ans {
			ans = val
		}
	}

	return ans
}
