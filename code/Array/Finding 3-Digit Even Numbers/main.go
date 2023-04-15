// Author 		: Praveen
// Date   		: 15/04/2023
// Question 	: https://leetcode.com/problems/finding-3-digit-even-numbers/
// Submission 	: https://leetcode.com/problems/finding-3-digit-even-numbers/submissions/934174808/

package main

import (
	"fmt"
	"sort"
)

func main() {
	digits := []int{3, 7, 5}
	fmt.Println(findEvenNumbers(digits))
}

func findEvenNumbers(digits []int) []int {
	var res []int

	m := make(map[int]struct{})

	for i, digit := range digits {
		if digit%2 == 0 {
			newArr := []int{}
			newArr = append(newArr, digits[:i]...)
			newArr = append(newArr, digits[i+1:]...)
			l := len(newArr)

			for j := 0; j < l; j++ {
				if newArr[j] != 0 {
					for k := 0; k < l; k++ {
						if j != k {
							n1 := newArr[j]
							n2 := newArr[k]
							sum := (n1 * 100) + (n2 * 10) + digit
							if _, ok := m[sum]; !ok {
								m[sum] = struct{}{}
								res = append(res, sum)
							}
						}

					}
				}

			}
		}
	}
	sort.Ints(res)
	return res
}
