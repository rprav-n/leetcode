// Author 		: Praveen
// Date   		: 29/04/2023
// Question 	: https://leetcode.com/problems/destination-city/
// Submission 	: https://leetcode.com/problems/destination-city/submissions/941719238/

package main

import "fmt"

func main() {
	paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	fmt.Println(destCity(paths))
}

func destCity(paths [][]string) string {
	var ans string

	start := make(map[string]struct{})
	end := make(map[string]struct{})

	for _, path := range paths {

		start[path[0]] = struct{}{}
		end[path[1]] = struct{}{}

	}

	for k := range end {
		if _, ok := start[k]; !ok {
			ans = k
			return ans
		}
	}

	return ans
}
