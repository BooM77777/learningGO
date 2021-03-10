package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	retVal := make([]int, 0, 10000)
	for iter := head; iter != nil; iter = iter.Next {
		retVal = append(retVal, iter.Val)
	}
	for l, r := 0, len(retVal)-1; l < r; l, r = l+1, r-1 {
		retVal[l], retVal[r] = retVal[r], retVal[l]
	}
	return retVal
}

func main() {

}
