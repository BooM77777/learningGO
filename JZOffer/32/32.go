package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	q := []*TreeNode{root}
	retVal := make([]int, 0, 1000)

	var (
		tmp  []*TreeNode
		head *TreeNode
	)

	for len(q) > 0 {
		tmp = make([]*TreeNode, 0, len(q)*2)
		for len(q) > 0 {
			head, q = q[0], q[1:]
			retVal = append(retVal, head.Val)
			if head.Left != nil {
				tmp = append(tmp, head.Left)
			}
			if head.Right != nil {
				tmp = append(tmp, head.Right)
			}
		}
		q = tmp
		// fmt.Println(len(q))
	}

	return retVal
}

func main() {

}
