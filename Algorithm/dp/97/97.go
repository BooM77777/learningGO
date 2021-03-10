package main

// TODO: 滚动数组优化
func isInterleave(s1 string, s2 string, s3 string) bool {

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	M, N := len(s1), len(s2)
	dp := make([][]bool, M+1)
	for i := range dp {
		dp[i] = make([]bool, N+1)
	}
	dp[0][0] = true

	for i := 0; i <= M; i++ {
		for j := 0; j <= N; j++ {
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && (s3[i+j-1] == s1[i-1]))
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && (s3[i+j-1] == s2[j-1]))
			}
		}
	}

	return dp[M][N]
}

func main() {

}
