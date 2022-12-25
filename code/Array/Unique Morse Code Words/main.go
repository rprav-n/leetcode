// Author 		: Praveen
// Date   		: 25/12/2022
// Question 	: https://leetcode.com/problems/unique-morse-code-words/
// Submission 	: https://leetcode.com/problems/unique-morse-code-words/submissions/865102564/

/*
	Question:
	--------------------------------
	International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows:

	'a' maps to ".-",
	'b' maps to "-...",
	'c' maps to "-.-.", and so on.

	For convenience, the full table for the 26 letters of the English alphabet is given below:

	[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]

	Given an array of strings words where each word can be written as a concatenation of the Morse code of each letter.

	For example, "cab" can be written as "-.-..--...", which is the concatenation of "-.-.", ".-", and "-...". We will call such a concatenation the transformation of a word.
	Return the number of different transformations among all words we have.


	Example 1:
	Input: words = ["gin","zen","gig","msg"]
	Output: 2
	Explanation: The transformation of each word is:
	"gin" -> "--...-."
	"zen" -> "--...-."
	"gig" -> "--...--."
	"msg" -> "--...--."
	There are 2 different transformations: "--...-." and "--...--.".


*/

package main

import "fmt"

func main() {
	words := []string{"a"}
	fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) int {

	m := make(map[rune]string)
	r := make(map[string]struct{})
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	table := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	for i, c := range alphabets {
		m[c] = table[i]
	}

	for _, word := range words {
		var rs string
		for _, c := range word {
			rs += m[c]
		}
		r[rs] = struct{}{}
	}

	return len(r)
}
