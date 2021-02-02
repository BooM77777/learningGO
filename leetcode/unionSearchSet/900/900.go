package main

func numIslands(grid [][]byte) (ret int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}
	M, N := len(grid), len(grid[0])
	var i, j int
	root := make([]int, M*N)
	for i = 0; i < M*N; i++ {
		root[i] = -1
	}
	hashTheNode := func(x, y int) int {
		return x*N + y
	}
	var findRoot func(int) int
	findRoot = func(node int) int {
		if root[node] != -1 {
			r := findRoot(root[node])
			root[node] = r
			return r
		}
		return node
	}
	union := func(n1, n2 int) {
		r1, r2 := findRoot(n1), findRoot(n2)
		if r1 != r2 {
			root[r1] = r2
		}
	}
	for i = 0; i < M; i++ {
		for j = 0; j < N; j++ {
			if grid[i][j] == '1' {
				if i+1 < M && grid[i+1][j] == '1' {
					union(hashTheNode(i, j), hashTheNode(i+1, j))
				}
				if j+1 < N && grid[i][j+1] == '1' {
					union(hashTheNode(i, j), hashTheNode(i, j+1))
				}
			}
		}
	}
	for i = 0; i < M; i++ {
		for j = 0; j < N; j++ {
			if root[hashTheNode(i, j)] == -1 && grid[i][j] == '1' {
				ret++
			}
		}
	}
	return
}

func main() {

}
