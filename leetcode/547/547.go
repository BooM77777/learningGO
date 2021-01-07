package main

import (
	"fmt"
)

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
}

// func bfs(isConnected [][]int, begin int, flag []bool) []bool {
// 	counter := len(isConnected)
// 	q := []int{begin}
// 	for len(q) > 0 {
// 		tmp := q[0]
// 		q = q[1:]
// 		for i := 0; i < counter; i++ {
// 			if isConnected[tmp][i] == 1 && !flag[i] {
// 				q = append(q, i)
// 				flag[i] = true
// 			}
// 		}
// 	}
// 	return flag
// }

// func findCircleNumBFS(isConnected [][]int) (ret int) {
// 	counter := len(isConnected)
// 	flag := make([]bool, counter)
// 	for i := 0; i < counter; i++ {
// 		if !flag[i] {
// 			flag = bfs(isConnected, i, flag)
// 			ret++
// 		}
// 	}
// 	return
// }

func findCircleNum(isConnected [][]int) (ret int) {

	counter := len(isConnected)
	root := make([]int, counter)
	for i := 0; i < counter; i++ {
		root[i] = -1
	}

	var findRoot func(int) int
	findRoot = func(i int) int {
		r := root[i]
		if r >= 0 {
			r = findRoot(i)
			return r
		}
		return i
	}

	merage := func(i, j int) {
		rootI, rootJ := findRoot(i), findRoot(j)
		if rootI != rootJ {
			if rootI < rootJ {
				root[rootI], root[rootJ] = root[rootI]+root[rootJ], rootI
			} else {
				root[rootJ], root[rootI] = root[rootI]+root[rootJ], rootJ
			}
		}
	}

	// find := func(i, j int) bool{
	//     rootI, rootJ := findRoot(i), findRoot(j)
	//     if rootI == rootJ {
	//         return true
	//     }
	//     return false
	// }

	for i := 0; i < counter; i++ {
		for j := i + 1; j < counter; j++ {
			if isConnected[i][j] == 1 {
				merage(i, j)
			}
		}
	}

	for i := 0; i < counter; i++ {
		if root[i] < 0 {
			ret++
		}
	}
	return
}
