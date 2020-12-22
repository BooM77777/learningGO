package main

import (
	"fmt"
)

type (
	//TreeNode ...
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	//Queue ...
	Queue struct {
		front *queueNode
		back  *queueNode
		size  int
	}

	//queueNode ...
	queueNode struct {
		prev *queueNode
		next *queueNode
		val  *TreeNode
	}
)

func New() *Queue {
	return &Queue{nil, nil, 0}
}

func (q *Queue) push(node *TreeNode) {
	qNode := &queueNode{nil, nil, node}
	if q.size == 0 {
		q.front = qNode
		q.back = qNode
	} else {
		q.back.next = qNode
		qNode.prev = q.back
		q.back = qNode
	}
	q.size++
}

func (q *Queue) pop() (*TreeNode, bool) {
	if q.size == 0 {
		return nil, false
	}
	ret := q.front
	q.front = q.front.next
	q.size--
	return ret.val, true
}

func (q *Queue) empty() bool {
	return q.size == 0
}

func main() {
	q := New()
	for i := 0; i < 11; i++ {
		tmpNode := &TreeNode{i, nil, nil}
		q.push(tmpNode)
	}
	for {
		if q.empty() {
			break
		}
		tmpNode, flag := q.pop()
		if flag {
			fmt.Println(tmpNode.Val)
		} else {
			fmt.Println("ERROR!")
		}
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ret [][]int
	q := New()
	cnt := 1
	q.push(root)
	for {
		if q.empty() {
			break
		}
		nextCnt := 0
		var subRet []int
		for i := 0; i < cnt; i++ {
			tmpNode, flag := q.pop()
			if !flag {
				return nil
			}
			if tmpNode == nil {
				continue
			}
			subRet = append(subRet, tmpNode.Val)
			q.push(tmpNode.Left)
			q.push(tmpNode.Right)
			nextCnt = nextCnt + 2
		}
		cnt = nextCnt
		if len(subRet) > 0 {
			ret = append(ret, subRet)
		}
	}
	for i, subRet := range ret {
		if i%2 == 1 {
			for i, j := 0, len(subRet)-1; i < j; i, j = i+1, j-1 {
				subRet[i], subRet[j] = subRet[j], subRet[i]
			}
		}
	}
	return ret
}
