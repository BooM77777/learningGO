package main

import "math"

/**
 * Definition for a binary tree node.
 *
 */

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func getMinimumDifference(root *TreeNode) int {
	retVal, pre := math.MaxInt64, -10086
	var midOrder func(*TreeNode)
	midOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		midOrder(root.Left)
		retVal = min(retVal, root.Val-pre)
		pre = root.Val
		midOrder(root.Right)
	}
	midOrder(root)
	return retVal
}

func main() {

}
