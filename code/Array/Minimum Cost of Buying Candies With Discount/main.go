// Author 		: Praveen
// Date   		: 17/01/2023
// Question 	: https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/
// Submission 	: https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/submissions/880046318/

package main

import (
	"fmt"
	"sort"
)

func main() {
	cost := []int{6, 5, 7, 9, 2, 2}
	fmt.Println(minimumCost(cost))

}

func minimumCost(cost []int) int {
	var ans int
	l := len(cost)

	sort.Ints(cost)
	fmt.Println(cost)

	i := l - 1
	c := 0
	for i >= 0 {
		if c > 2 {
			c = 0
		}
		if c != 2 {
			fmt.Println(cost[i], c)
			ans += cost[i]
		}
		c++
		i--
	}

	return ans
}
