package main

func min(a ...int) (ret int) {
	ret = a[0]
	for i := 1; i < len(a); i++ {
		if ret > a[i] {
			ret = a[i]
		}
	}
	return
}

func maximalSquare(matrix [][]byte) (ret int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	M, N, i, j := len(matrix), len(matrix[0]), 0, 0
	dp := make([][]int, M+1)
	for i = 0; i < M+1; i++ {
		dp[i] = make([]int, N+1)
	}
	for i = 0; i < M; i++ {
		for j = 0; j < N; j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(dp[i][j], dp[i+1][j], dp[i][j+1]) + 1
				if ret < dp[i+1][j+1] {
					ret = dp[i+1][j+1]
				}
			}
		}
	}
	ret = ret * ret
	return
}

func main() {

}
