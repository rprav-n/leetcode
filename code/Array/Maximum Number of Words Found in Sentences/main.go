// Author 		: Praveen
// Date   		: 04/12/2022
// Question 	: https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
// Submission 	: https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/submissions/854566652/

/*
	Question:
	--------------------------------

	A sentence is a list of words that are separated by a single space with no leading or trailing spaces.

	You are given an array of strings sentences, where each sentences[i] represents a single sentence.

	Return the maximum number of words that appear in a single sentence.


	Example 1:

	Input: sentences = ["alice and bob love leetcode", "i think so too", "this is great thanks very much"]
	Output: 6
	Explanation:
	- The first sentence, "alice and bob love leetcode", has 5 words in total.
	- The second sentence, "i think so too", has 4 words in total.
	- The third sentence, "this is great thanks very much", has 6 words in total.
	Thus, the maximum number of words in a single sentence comes from the third sentence, which has 6 words.


	Notes:
	--------------------------------
	1. Create a variable called "maxNumOfWords" of type int which is the result
	2. Loop through the sentences array
		2.1 Use strings std lib to split the sentence with sep(separator) as " " which spits out slice of strings
		2.2 Get the lenth (l) of the string slice
		2.3 Check if length(l) is greater than maxNumOfWords
			If so, assign maxNumOfWords = length(l)
	3. Return the maxNumOfWords

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}

func mostWordsFound(sentences []string) int {
	var maxNumOfWords int

	for _, sentence := range sentences {
		sentenceArr := strings.Split(sentence, " ")
		l := len(sentenceArr)
		if l > maxNumOfWords {
			maxNumOfWords = l
		}
	}

	return maxNumOfWords
}
