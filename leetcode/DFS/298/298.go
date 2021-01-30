package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func longestConsecutive(root *TreeNode) (ret int) {

	var dfs func(root *TreeNode, dist int)
	dfs = func(root *TreeNode, dist int) {
		if root == nil {
			return
		}
		ret = max(ret, dist)
		if root.Left != nil {
			if root.Val+1 == root.Left.Val {
				dfs(root.Left, dist+1)
			} else {
				dfs(root.Left, 1)
			}
		}
		if root.Right != nil {
			if root.Val+1 == root.Right.Val {
				dfs(root.Right, dist+1)
			} else {
				dfs(root.Right, 1)
			}
		}
	}
	dfs(root, 1)
	return
}

func main() {}
