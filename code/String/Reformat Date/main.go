// Author 		: Praveen
// Date   		: 24/05/2023
// Question 	: https://leetcode.com/problems/reformat-date/
// Submission 	: https://leetcode.com/problems/reformat-date/submissions/956571791/

package main

import (
	"fmt"
	"strings"
)

func main() {
	date := "6th Jun 1933"
	fmt.Println(reformatDate(date))
}

func reformatDate(date string) string {
	var ans string

	monthMap := make(map[string]string)
	monthMap["Jan"] = "01"
	monthMap["Feb"] = "02"
	monthMap["Mar"] = "03"
	monthMap["Apr"] = "04"
	monthMap["May"] = "05"
	monthMap["Jun"] = "06"
	monthMap["Jul"] = "07"
	monthMap["Aug"] = "08"
	monthMap["Sep"] = "09"
	monthMap["Oct"] = "10"
	monthMap["Nov"] = "11"
	monthMap["Dec"] = "12"

	dates := strings.Split(date, " ")

	l := len(dates)

	for i := l - 1; i >= 0; i-- {
		d := dates[i]
		if i == 0 {
			ds := strings.Replace(d, "st", "", 1)
			ds = strings.Replace(ds, "nd", "", 1)
			ds = strings.Replace(ds, "rd", "", 1)
			ds = strings.Replace(ds, "th", "", 1)
			if len(ds) == 1 {
				ans += "0" + ds
			} else {
				ans += ds
			}

		} else if i == 1 {
			t := monthMap[d]
			ans += fmt.Sprintf("%s-", t)
		} else {
			ans += d + "-"
		}
	}

	return ans
}
