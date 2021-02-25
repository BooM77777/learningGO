package main

import "sort"

// 边
type Edge struct {
	n1, n2 int
	weight int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minimumEffortPath(heights [][]int) int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return -1
	}
	M, N := len(heights), len(heights[0])
	if M*N == 1 {
		return 0
	}
	edges := make([]Edge, 0, M*N)
	getNID := func(x, y int) int {
		return x*N + y
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if i < M-1 {
				edges = append(edges, Edge{getNID(i, j), getNID(i+1, j), abs(heights[i+1][j] - heights[i][j])})
			}
			if j < N-1 {
				edges = append(edges, Edge{getNID(i, j), getNID(i, j+1), abs(heights[i][j+1] - heights[i][j])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight })
	root := make([]int, M*N)
	for i := range root {
		root[i] = i
	}
	var findRoot func(n int) int
	findRoot = func(n int) int {
		if root[n] != n {
			r := findRoot(root[n])
			root[n] = r
			return r
		}
		return n
	}
	union := func(n1, n2 int) {
		r1, r2 := findRoot(n1), findRoot(n2)
		if r1 != r2 {
			root[r1] = r2
		}
	}
	bN, eN := getNID(0, 0), getNID(M-1, N-1)
	for _, edge := range edges {
		union(edge.n1, edge.n2)
		// fmt.Println(root)
		if findRoot(bN) == findRoot(eN) {
			return edge.weight
		}
	}
	return -1
}

func main() {

}
