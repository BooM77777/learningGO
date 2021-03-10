package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	M, N := len(matrix), len(matrix[0])
	x, y := 0, N-1
	for x < M && y >= 0 {
		if matrix[x][y] < target {
			x++
		} else if matrix[x][y] > target {
			y--
		} else {
			return true
		}
	}
	return false
}

func main() {

}
