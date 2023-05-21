// Author 		: Praveen
// Date   		: 21/05/2023
// Question 	: https://leetcode.com/problems/minimum-number-of-operations-to-convert-time/
// Submission 	: https://leetcode.com/problems/minimum-number-of-operations-to-convert-time/submissions/954654184/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	current := "02:30"
	correct := "04:35"

	fmt.Println(convertTime(current, correct))
}

func convertTime(current string, correct string) int {
	var count int

	currentArr := strings.Split(current, ":")
	correctArr := strings.Split(correct, ":")

	currentMins := 0
	correctMins := 0

	for i, s := range currentArr {
		si, _ := strconv.Atoi(s)
		if i == 0 {
			currentMins += si * 60
		} else {
			currentMins += si
		}
	}

	for i, s := range correctArr {
		si, _ := strconv.Atoi(s)
		if i == 0 {
			correctMins += si * 60
		} else {
			correctMins += si
		}
	}

	for currentMins != correctMins {
		sub := correctMins - currentMins
		if sub >= 60 {
			currentMins += 60
		} else if sub >= 15 {
			currentMins += 15
		} else if sub >= 5 {
			currentMins += 5
		} else {
			currentMins += 1
		}
		count++
	}

	return count
}
