// Author 		: Praveen
// Date   		: 12/12/2022
// Question 	: https://leetcode.com/problems/count-items-matching-a-rule/
// Submission 	: https://leetcode.com/problems/count-items-matching-a-rule/submissions/858736720/

/*
	Question:
	--------------------------------
	You are given an array items, where each items[i] = [typei, colori, namei] describes the type, color, and name of the ith item. You are also given a rule represented by two strings, ruleKey and ruleValue.

	The ith item is said to match the rule if one of the following is true:

	ruleKey == "type" and ruleValue == typei.
	ruleKey == "color" and ruleValue == colori.
	ruleKey == "name" and ruleValue == namei.
	Return the number of items that match the given rule.


	Example 1:
	Input: items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]], ruleKey = "color", ruleValue = "silver"
	Output: 1
	Explanation: There is only one item matching the given rule, which is ["computer","silver","lenovo"].


	Example 2:
	Input: items = [["phone","blue","pixel"],["computer","silver","phone"],["phone","gold","iphone"]], ruleKey = "type", ruleValue = "phone"
	Output: 2
	Explanation: There are only two items matching the given rule, which are ["phone","blue","pixel"] and ["phone","gold","iphone"]. Note that the item ["computer","silver","phone"] does not match.


	Notes:
	--------------------------------


*/

package main

import "fmt"

func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}
	ruleKey := "type"
	ruleValue := "phone"

	fmt.Println(countMatches(items, ruleKey, ruleValue))
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var count, index int

	if ruleKey == "type" {
		index = 0
	} else if ruleKey == "color" {
		index = 1
	} else {
		index = 2
	}

	for _, item := range items {
	innerLoop:
		for i, value := range item {
			if i == index && value == ruleValue {
				count++
				break innerLoop
			}
		}
	}

	return count
}
