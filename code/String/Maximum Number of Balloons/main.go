// Author 		: Praveen
// Date   		: 07/06/2023
// Question 	: https://leetcode.com/problems/maximum-number-of-balloons/
// Submission 	: https://leetcode.com/problems/maximum-number-of-balloons/submissions/966037938/

package main

import (
	"fmt"
	"math"
)

func main() {
	text := "krhizmmgmcrecekgyljqkldocicziihtgpqwbticmvuyznragqoyrukzopfmjhjjxemsxmrsxuqmnkrzhgvtgdgtykhcglurvppvcwhrhrjoislonvvglhdciilduvuiebmffaagxerjeewmtcwmhmtwlxtvlbocczlrppmpjbpnifqtlninyzjtmazxdbzwxthpvrfulvrspycqcghuopjirzoeuqhetnbrcdakilzmklxwudxxhwilasbjjhhfgghogqoofsufysmcqeilaivtmfziumjloewbkjvaahsaaggteppqyuoylgpbdwqubaalfwcqrjeycjbbpifjbpigjdnnswocusuprydgrtxuaojeriigwumlovafxnpibjopjfqzrwemoinmptxddgcszmfprdrichjeqcvikynzigleaajcysusqasqadjemgnyvmzmbcfrttrzonwafrnedglhpudovigwvpimttiketopkvqw"
	fmt.Println(maxNumberOfBalloons(text))
}

func maxNumberOfBalloons(text string) int {
	count := math.MaxInt

	word := "balon"

	m := make(map[rune]int)

	for _, c := range word {
		m[c] = 0
	}

	for _, c := range text {
		m[c] += 1
	}

	m['l'] /= 2
	m['o'] /= 2

	for _, c := range word {
		v := m[c]
		if v < count {
			count = v
		}
	}

	return count

}
