package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	zeroNode := &ListNode{Val: 0, Next: head}

	tail := zeroNode
	for ; tail.Next != nil; tail = tail.Next {
	}

	iter := zeroNode.Next
	for iter != tail {
		tmpNode := iter.Next
		tail.Next, iter.Next = iter, tail.Next
		iter = tmpNode
	}
	return tail
}

func main() {

}
