package main

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	M, N, i, j := len(matrix), len(matrix[0]), 0, 0
	sum := make([][]int, M+1)
	for i = 0; i < M+1; i++ {
		sum[i] = make([]int, N+1)
	}
	for i = 0; i < M; i++ {
		for j = 0; j < N; j++ {
			sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{
		sum: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row1][col2+1] - this.sum[row2+1][col1] + this.sum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {

}
