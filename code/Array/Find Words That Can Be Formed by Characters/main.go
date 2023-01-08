// Author 		: Praveen
// Date   		: 08/01/2023
// Question 	: https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
// Submission 	: https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/submissions/874037255/

/*
	Question:
	--------------------------------
	You are given an array of strings words and a string chars.

	A string is good if it can be formed by characters from chars (each character can only be used once).

	Return the sum of lengths of all good strings in words.


	Example 1:
	Input: words = ["cat","bt","hat","tree"], chars = "atach"
	Output: 6
	Explanation: The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.


*/

package main

import (
	"fmt"
)

func main() {
	words := []string{"dyiclysmffuhibgfvapygkorkqllqlvokosagyelotobicwcmebnpznjbirzrzsrtzjxhsfpiwyfhzyonmuabtlwin", "ndqeyhhcquplmznwslewjzuyfgklssvkqxmqjpwhrshycmvrb", "ulrrbpspyudncdlbkxkrqpivfftrggemkpyjl", "boygirdlggnh", "xmqohbyqwagkjzpyawsydmdaattthmuvjbzwpyopyafphx", "nulvimegcsiwvhwuiyednoxpugfeimnnyeoczuzxgxbqjvegcxeqnjbwnbvowastqhojepisusvsidhqmszbrnynkyop", "hiefuovybkpgzygprmndrkyspoiyapdwkxebgsmodhzpx", "juldqdzeskpffaoqcyyxiqqowsalqumddcufhouhrskozhlmobiwzxnhdkidr", "lnnvsdcrvzfmrvurucrzlfyigcycffpiuoo", "oxgaskztzroxuntiwlfyufddl", "tfspedteabxatkaypitjfkhkkigdwdkctqbczcugripkgcyfezpuklfqfcsccboarbfbjfrkxp", "qnagrpfzlyrouolqquytwnwnsqnmuzphne", "eeilfdaookieawrrbvtnqfzcricvhpiv", "sisvsjzyrbdsjcwwygdnxcjhzhsxhpceqz", "yhouqhjevqxtecomahbwoptzlkyvjexhzcbccusbjjdgcfzlkoqwiwue", "hwxxighzvceaplsycajkhynkhzkwkouszwaiuzqcleyflqrxgjsvlegvupzqijbornbfwpefhxekgpuvgiyeudhncv", "cpwcjwgbcquirnsazumgjjcltitmeyfaudbnbqhflvecjsupjmgwfbjo", "teyygdmmyadppuopvqdodaczob", "qaeowuwqsqffvibrtxnjnzvzuuonrkwpysyxvkijemmpdmtnqxwekbpfzs", "qqxpxpmemkldghbmbyxpkwgkaykaerhmwwjonrhcsubchs"}
	chars := "usdruypficfbpfbivlrhutcgvyjenlxzeovdyjtgvvfdjzcmikjraspdfp"
	fmt.Println(countCharacters(words, chars))
}

func countCharacters(words []string, chars string) int {
	var ans int

	m := make(map[rune]int)
	for _, c := range chars {
		if _, ok := m[c]; ok {
			m[c] = m[c] + 1
		} else {
			m[c] = 1
		}
	}

loop:
	for _, word := range words {

		wm := make(map[rune]int)
		for _, c := range word {
			if _, ok := wm[c]; ok {
				wm[c] = wm[c] + 1
			} else {
				wm[c] = 1
			}
		}

		for k, v := range wm {
			if _, ok := m[k]; !ok {
				continue loop
			} else if v > m[k] {
				continue loop
			}
		}

		ans += len(word)
	}

	return ans
}
