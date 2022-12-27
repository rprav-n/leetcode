// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/
// Submission 	: https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/submissions/866466688/

/*
	Question:
	--------------------------------
	Given two integer arrays startTime and endTime and given an integer queryTime.

	The ith student started doing their homework at the time startTime[i] and finished it at time endTime[i].

	Return the number of students doing their homework at time queryTime. More formally, return the number of students where queryTime lays in the interval [startTime[i], endTime[i]] inclusive.


	Example 1:
	Input: startTime = [1,2,3], endTime = [3,2,7], queryTime = 4
	Output: 1
	Explanation: We have 3 students where:
	The first student started doing homework at time 1 and finished at time 3 and wasn't doing anything at time 4.
	The second student started doing homework at time 2 and finished at time 2 and also wasn't doing anything at time 4.
	The third student started doing homework at time 3 and finished at time 7 and was the only student doing homework at time 4.


*/

package main

import "fmt"

func main() {
	startTime := []int{4}
	endTime := []int{4}
	queryTime := 4
	fmt.Println(busyStudent(startTime, endTime, queryTime))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var result int

	l := len(startTime)

	for i := 0; i < l; i++ {
		s := startTime[i]
		e := endTime[i]
		if queryTime >= s && queryTime <= e {
			result++
		}
	}

	return result
}
