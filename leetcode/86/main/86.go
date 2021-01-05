package main

import "fmt"

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := []int{1, 4, 3, 2, 5, 2}
	x := 3
	head := &ListNode{0, nil}
	tmp := head
	for _, item := range l {
		tmp.Next = &ListNode{item, nil}
		tmp = tmp.Next
	}
	head = head.Next
	print(head)
	partition(head, x)
	print(head)
}

func print(head *ListNode) {
	tmpNode := head
	cnt := 0
	for tmpNode != nil {
		cnt++
		fmt.Printf("%d\t", tmpNode.Val)
		tmpNode = tmpNode.Next
		if cnt == 10 {
			break
		}
	}
	fmt.Println()
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	zeroNode := &ListNode{0, head}
	pNode, tmpIter := zeroNode, head
	for tmpIter != nil {
		pNode, tmpIter = tmpIter, tmpIter.Next
	}
	smallerIter, biggerIter, lastOne := zeroNode, pNode, pNode
	pNode, tmpIter = zeroNode, head

	isTheLastOne := false
	for tmpIter != nil {
		isTheLastOne = (tmpIter == lastOne)
		if tmpIter.Val < x {
			// 插入smallerIter
			pNode.Next = tmpIter.Next
			smallerIter.Next, tmpIter.Next = tmpIter.Next, smallerIter.Next
		} else {
			// 插入结尾
			pNode.Next, biggerIter.Next = tmpIter.Next, tmpIter
		}
		if isTheLastOne {
			break
		}
		tmpIter = pNode.Next
	}
	return zeroNode.Next
}
