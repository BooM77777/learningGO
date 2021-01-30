package main

import "fmt"

func main() {
	board := [][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}, []int{0, 0, 0}}
	gameOfLife(board)
	fmt.Println(board)
}

const (
	liveThenLive = 1
	deadThenDead = 0
	liveThenDead = -1
	deadThenLive = -2
)

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	M, N := len(board), len(board[0])

	getState := func(i, j int) int {
		switch board[i][j] {
		case liveThenLive:
			return 1
		case liveThenDead:
			return 1
		case deadThenDead:
			return 0
		case deadThenLive:
			return 0
		}
		return -1
	}
	count := func(i, j int) (ret int) {
		if i > 0 {
			ret += getState(i-1, j)
			if j > 0 {
				ret += getState(i-1, j-1)
			}
			if j < N-1 {
				ret += getState(i-1, j+1)
			}
		}
		if i < M-1 {
			ret += getState(i+1, j)
			if j > 0 {
				ret += getState(i+1, j-1)
			}
			if j < N-1 {
				ret += getState(i+1, j+1)
			}
		}
		if j > 0 {
			ret += getState(i, j-1)
		}
		if j < N-1 {
			ret += getState(i, j+1)
		}
		return
	}

	var state int
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			state = count(i, j)
			if state < 2 {
				// 小于2个
				if board[i][j] == 1 {
					board[i][j] = liveThenDead
				} else {
					board[i][j] = deadThenDead
				}
			} else if state <= 3 {
				// 2个或3个
				if board[i][j] == 1 {
					board[i][j] = liveThenLive
				} else {
					if state == 3 {
						// 复活
						board[i][j] = deadThenLive
					} else {
						board[i][j] = deadThenDead
					}
				}
			} else {
				// 4个及以上
				if board[i][j] == 1 {
					board[i][j] = liveThenDead
				} else {
					board[i][j] = deadThenDead
				}
			}
		}
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			switch board[i][j] {
			case liveThenDead:
				board[i][j] = 0
				break
			case deadThenLive:
				board[i][j] = 1
				break
			case liveThenLive:
				board[i][j] = 1
				break
			case deadThenDead:
				board[i][j] = 0
				break
			default:
				board[i][j] = -1
			}
		}
	}
}
