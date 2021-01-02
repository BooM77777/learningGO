package main

import (
	"fmt"
)

const (
	//RED ...
	RED = iota
	// BLACK ...
	BLACK
)

type (
	//RedBlackTree ...
	RedBlackTree struct {
		root *TreeNode
	}
	//TreeNode ...
	TreeNode struct {
		left   *TreeNode
		right  *TreeNode
		parent *TreeNode
		data   int
		color  int
	}
)

//Insert 红黑树的插入
func (t *RedBlackTree) Insert(data int) (err bool) {

	// 默认红色，插入红色节点不会影响红黑树的黑平衡
	node := &TreeNode{nil, nil, nil, data, RED}
	// 查找到对应的插入位置
	var pNode *TreeNode
	tmpNode := t.root
	for tmpNode != nil {
		pNode = tmpNode
		if data < pNode.data {
			tmpNode = tmpNode.left
		} else if data > pNode.data {
			tmpNode = tmpNode.right
		} else {
			return false
		}
	}
	if pNode == nil {
		// 空树，直接插入
		t.root = node
	} else {
		if data < pNode.data {
			pNode.left = node
			node.parent = pNode
		} else if data > pNode.data {
			pNode.right = node
			node.parent = pNode
		} else {
			return false
		}
	}

	inserFixedUp(t, node)

	return true
}

func inserFixedUp(tree *RedBlackTree, node *TreeNode) {
	for node.parent != nil && node.parent.color == RED {
		if node.parent.parent.left == node.parent {
			// 如果是左节点
			// 定义祖父节点、父节点、叔节点
			grandpaNode, pNode, uncleNode := node.parent.parent, node.parent, node.parent.parent.right
			if uncleNode != nil && uncleNode.color == RED {
				grandpaNode.color = RED
				pNode.color, uncleNode.color = BLACK, BLACK
				node = grandpaNode
			} else {
				if pNode == grandpaNode.right {
					node = pNode
					leftRotate(tree, node)
				}
				node.parent.parent.color = RED
				node.parent.color = BLACK
				node.color = RED
				rightRotate(tree, grandpaNode)
			}
		} else {
			// 如果是右节点
			grandpaNode, pNode, uncleNode := node.parent.parent, node.parent, node.parent.parent.left
			if uncleNode != nil && uncleNode.color == RED {
				grandpaNode.color = RED
				pNode.color, uncleNode.color = BLACK, BLACK
				node = grandpaNode
			} else {
				if node == pNode.left {
					node = pNode
					rightRotate(tree, node)
				}
				node.parent.parent.color = RED
				node.parent.color = BLACK
				node.color = RED
				leftRotate(tree, grandpaNode)
			}
		}
	}
	tree.root.color = BLACK
}

//leftRotate: 以某个结点作为支点(旋转结点)，其右子结点变为旋转结点的父结点，右子结点的左子结点变为旋转结点的右子结点，左子结点保持不变
func leftRotate(tree *RedBlackTree, node *TreeNode) {
	// pNode := node.parent

	rightChild := node.right
	// 右子结点的左子结点变为旋转结点的右子结点
	node.right = rightChild.left
	// 其右子结点变为旋转结点的父结点
	node.parent, rightChild.parent, rightChild.left = rightChild, node.parent, node
	if rightChild.parent == nil {
		tree.root = rightChild
	} else {
		if node == rightChild.parent.left {
			rightChild.parent.left = rightChild
		} else {
			rightChild.parent.right = rightChild
		}
	}
}

//rightRotate: 以某个结点作为支点(旋转结点)，其左子结点变为旋转结点的父结点，左子结点的右子结点变为旋转结点的左子结点，右子结点保持不变。
func rightRotate(tree *RedBlackTree, node *TreeNode) {

	leftChild := node.left
	node.left = leftChild.right
	node.parent, leftChild.parent, leftChild.right = leftChild, node.parent, node
	if leftChild.parent == nil {
		tree.root = leftChild
	} else {
		if node == leftChild.parent.left {
			leftChild.parent.left = leftChild
		} else {
			leftChild.parent.right = leftChild
		}
	}
}

//PreOrder 红黑树的前序遍历
func PreOrder(tree *RedBlackTree) {
	fmt.Println("先序遍历 : ")
	preOrder(tree.root)
	fmt.Println()
}

//MidOrder 二叉树的中序遍历
func MidOrder(tree *RedBlackTree) {
	fmt.Println("中序遍历 : ")
	midOrder(tree.root)
	fmt.Println()
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	color := "black"
	if root.color == RED {
		color = "red"
	}
	fmt.Printf("<%c, %s>\t", root.data, color)
	preOrder(root.left)
	preOrder(root.right)
}

func midOrder(root *TreeNode) {
	if root == nil {
		return
	}
	color := "black"
	if root.color == RED {
		color = "red"
	}
	midOrder(root.left)
	fmt.Printf("<%c, %s>\t", root.data, color)
	midOrder(root.right)
}

func main() {
	tree := &RedBlackTree{nil}
	str := "ASERCHINGX"
	for _, c := range str {
		fmt.Printf("===> 插入%c\n", c)
		tree.Insert(int(c))
		PreOrder(tree)
		MidOrder(tree)
	}
}
