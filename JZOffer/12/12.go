package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	M, N := len(board), len(board[0])

	mark := make([][]bool, M)
	for i := range mark {
		mark[i] = make([]bool, N)
	}

	directions := [4][2]int{[2]int{1, 0}, [2]int{0, 1}, [2]int{0, -1}, [2]int{-1, 0}}

	var dfs func(x, y int, index int) bool
	dfs = func(x, y int, index int) bool {

		if board[x][y] != word[index] {
			return false
		}
		index++
		if index == len(word) {
			return true
		}
		mark[x][y] = true
		ret := false
		for _, dir := range directions {
			tmpX, tmpY := x+dir[0], y+dir[1]
			if tmpX < 0 || tmpX >= M || tmpY < 0 || tmpY >= N {
				continue
			} else if mark[tmpX][tmpY] {
				continue
			}
			ret = dfs(tmpX, tmpY, index)
			if ret {
				break
			}
		}
		mark[x][y] = false
		return ret
	}

	for i := range board {
		for j := range board[i] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {

}
