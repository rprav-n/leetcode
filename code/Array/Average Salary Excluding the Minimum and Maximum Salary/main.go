// Author 		: Praveen
// Date   		: 31/01/2022
// Question 	: https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
// Submission 	: https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/submissions/888918005/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	fmt.Println(average(salary))
}

func average(salary []int) float64 {
	var ans float64

	sort.Ints(salary)

	l := len(salary)

	for i := 1; i < l-1; i++ {
		ans += float64(salary[i])
	}
	ans = Round(ans/float64(l-2), 5)

	return ans
}

func Round(val float64, precision int) float64 {
	return math.Round(val*(math.Pow10(precision))) / math.Pow10(precision)
}
