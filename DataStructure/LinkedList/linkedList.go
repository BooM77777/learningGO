package main

import "fmt"

type (
	// LinkedListNode ...
	LinkedListNode struct {
		Val  int
		Next *LinkedListNode
	}

	// LinkedList ...
	LinkedList struct {
		Head *LinkedListNode
	}
)

// NewLinkedList ...
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: &LinkedListNode{
			Val:  0,
			Next: nil,
		},
	}
}

// InsertToHead ...
func (l *LinkedList) InsertToHead(data int) {
	tmpNode := &LinkedListNode{
		Val:  data,
		Next: nil,
	}
	l.Head.Next, tmpNode.Next = tmpNode, l.Head.Next
}

// Insert ...
func (l *LinkedList) Insert(data int) {
	var iter *LinkedListNode
	for iter = l.Head; iter.Next != nil; iter = iter.Next {
	}
	iter.Next = &LinkedListNode{
		Val:  data,
		Next: nil,
	}
}

// PopTail ...
func (l *LinkedList) PopTail() (int, bool) {
	if l.Head.Next == nil {
		return 0, false
	}
	var piter, iter *LinkedListNode
	for iter = l.Head; iter.Next != nil; piter, iter = iter, iter.Next {
	}
	piter.Next = nil
	return iter.Val, true
}

// PopHead ...
func (l *LinkedList) PopHead() (int, bool) {
	if l.Head.Next == nil {
		return 0, false
	}
	aimNode := l.Head.Next
	l.Head.Next = aimNode.Next
	return aimNode.Val, true
}

// Println ...
func (l *LinkedList) Println() {
	for iter := l.Head.Next; iter != nil; iter = iter.Next {
		fmt.Printf("%d ", iter.Val)
	}
	fmt.Println()
}

func main() {
	linkedList := NewLinkedList()
	list := []int{3, 6, 5, 7, 3, 4, 8, 7, 6}
	for _, num := range list {
		linkedList.InsertToHead(num)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(linkedList.PopTail())
	}
	linkedList.Println()
}
