package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	iter1, iter2 := headA, headB
	flag := 0
	for flag < 3 && iter1 != iter2 {
		iter1, iter2 = iter1.Next, iter2.Next
		if iter1 == nil {
			iter1 = headB
			flag++
		}
		if iter2 == nil {
			iter2 = headA
			flag++
		}
	}
	if flag < 3 {
		return iter1
	} else {
		return nil
	}
}

func main() {

}
