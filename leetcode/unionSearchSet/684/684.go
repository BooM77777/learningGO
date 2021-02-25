package main

func findRedundantConnection(edges [][]int) []int {
	N := len(edges)
	root := make([]int, N)
	for i := range root {
		root[i] = i
	}
	var findRoot func(int) int
	findRoot = func(node int) int {
		if root[node] != node {
			root[node] = findRoot(root[node])
		}
		return root[node]
	}
	merage := func(c1, c2 int) bool {
		r1, r2 := findRoot(c1), findRoot(c2)
		if r1 == r2 {
			return false
		}
		root[r1] = r2
		return true
	}
	for _, edge := range edges {
		if merage(edge[0]-1, edge[1]-1) == false {
			return edge
		}
	}
	return []int{}
}

func main() {

}
