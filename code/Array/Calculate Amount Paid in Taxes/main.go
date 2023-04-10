// Author 		: Praveen
// Date   		: 10/04/2023
// Question 	: https://leetcode.com/problems/calculate-amount-paid-in-taxes/
// Submission 	: https://leetcode.com/problems/calculate-amount-paid-in-taxes/submissions/931028417/

package main

import (
	"fmt"
	"math"
)

func main() {
	brackets := [][]int{{2, 50}}
	income := 0
	fmt.Println(calculateTax(brackets, income))
}

func calculateTax(brackets [][]int, income int) float64 {
	var res float64
	l := len(brackets)

	for i := 0; i < l; i++ {
		if i == 0 {
			val := brackets[i][0]
			if val > income {
				val = income
			}
			income -= val
			res += float64(val*brackets[i][1]) / 100
		} else {
			val := brackets[i][0] - brackets[i-1][0]
			if val > income {
				val = income
			}
			income -= val

			res += float64(val*brackets[i][1]) / 100

		}

	}
	res = toFixed(res, 5)
	return res
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
