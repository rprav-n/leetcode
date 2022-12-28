// Author 		: Praveen
// Date   		: 27/12/2022
// Question 	: https://leetcode.com/problems/merge-similar-items/
// Submission 	: https://leetcode.com/problems/merge-similar-items/submissions/866711379/

/*
	Question:
	--------------------------------
	You are given two 2D integer arrays, items1 and items2, representing two sets of items. Each array items has the following properties:

	items[i] = [valuei, weighti] where valuei represents the value and weighti represents the weight of the ith item.
	The value of each item in items is unique.
	Return a 2D integer array ret where ret[i] = [valuei, weighti], with weighti being the sum of weights of all items with value valuei.

	Note: ret should be returned in ascending order by value.


	Example 1:
	Input: items1 = [[1,1],[4,5],[3,8]], items2 = [[3,1],[1,5]]
	Output: [[1,6],[3,9],[4,5]]
	Explanation:
	The item with value = 1 occurs in items1 with weight = 1 and in items2 with weight = 5, total weight = 1 + 5 = 6.
	The item with value = 3 occurs in items1 with weight = 8 and in items2 with weight = 1, total weight = 8 + 1 = 9.
	The item with value = 4 occurs in items1 with weight = 5, total weight = 5.
	Therefore, we return [[1,6],[3,9],[4,5]].


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	items1 := [][]int{{5, 1}, {4, 2}, {3, 3}, {2, 4}, {1, 5}}
	items2 := [][]int{{7, 1}, {6, 2}, {5, 3}, {4, 4}}
	fmt.Println(mergeSimilarItems(items1, items2))
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var result [][]int
	m := make(map[int]struct{})

	for i := 0; i < len(items1); i++ {
		v := items1[i][0]
		w := items1[i][1]
		for j := 0; j < len(items2); j++ {
			v2 := items2[j][0]
			if v == v2 {
				w += items2[j][1]
			}
		}
		m[v] = struct{}{}
		result = append(result, []int{v, w})

	}
	for _, n := range items1 {
		if _, ok := m[n[0]]; !ok {
			m[n[0]] = struct{}{}
			result = append(result, n)
		}
	}
	for _, n := range items2 {
		if _, ok := m[n[0]]; !ok {
			m[n[0]] = struct{}{}
			result = append(result, n)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
