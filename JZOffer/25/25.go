package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	iter, iter1, iter2 := head, l1, l2
	for iter1 != nil && iter2 != nil {
		if iter1.Val <= iter2.Val {
			iter.Next, iter1 = iter1, iter1.Next
		} else {
			iter.Next, iter2 = iter2, iter2.Next
		}
		iter = iter.Next
	}
	for iter1 != nil {
		iter.Next = iter1
		iter, iter1 = iter.Next, iter1.Next
	}
	for iter2 != nil {
		iter.Next = iter2
		iter, iter2 = iter.Next, iter2.Next
	}
	return head.Next
}

func main() {

}
