// Author 		: Praveen
// Date   		: 21/01/2023
// Question 	: https://leetcode.com/problems/slowest-key/
// Submission 	: https://leetcode.com/problems/slowest-key/submissions/882331414/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	releaseTimes := []int{12, 23, 36, 46, 62}
	keysPressed := "spuda"
	fmt.Println(slowestKey(releaseTimes, keysPressed))
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var ans byte
	var arr []int

	m := map[string]int{}

	l := len(releaseTimes)

	for i := 0; i < l; i++ {
		if i == 0 {
			m[fmt.Sprintf("%d_%d", keysPressed[i], i)] = releaseTimes[i]
		} else {
			m[fmt.Sprintf("%d_%d", keysPressed[i], i)] = releaseTimes[i] - releaseTimes[i-1]
		}
	}

	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	for k, v := range m {
		if v == max {
			v := strings.Split(k, "_")
			vi, _ := strconv.Atoi(v[0])
			arr = append(arr, vi)
		}
	}

	ma := arr[0]
	for _, v := range arr {
		if v > ma {
			ma = v
		}
	}

	for _, c := range keysPressed {
		if int(c) == ma {
			ans = byte(c)
		}
	}

	return ans
}
