// Author 		: Praveen
// Date   		: 22/01/2023
// Question 	: https://leetcode.com/problems/design-an-ordered-stream/
// Submission 	: https://leetcode.com/problems/design-an-ordered-stream/submissions/883167704/

package main

func main() {
	os := Constructor(5)  // Inserts (3, "ccccc"), returns [].
	os.Insert(1, "aaaaa") // Inserts (1, "aaaaa"), returns ["aaaaa"].
	os.Insert(2, "bbbbb") // Inserts (2, "bbbbb"), returns ["bbbbb", "ccccc"].
	os.Insert(5, "eeeee") // Inserts (5, "eeeee"), returns [].
	os.Insert(4, "ddddd")
}

type OrderedStream struct {
	Arr []string
	Ptr int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		Arr: make([]string, n),
		Ptr: 0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {

	this.Arr[idKey-1] = value

	if this.Arr[this.Ptr] == "" {
		return []string{}
	} else {
		for i, v := range this.Arr[this.Ptr:] {
			if v == "" {
				temp := this.Ptr
				this.Ptr += i
				return this.Arr[temp:this.Ptr]
			}
		}
	}
	return this.Arr[this.Ptr:]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
