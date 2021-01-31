package main

func min(a ...int) int {
	ret := a[0]
	for _, d := range a {
		if ret > d {
			ret = d
		}
	}
	return ret
}

func minDistance(word1 string, word2 string) int {
	M, N := len(word1), len(word2)
	if M == 0 || N == 0 {
		return M | N
	}
	dp := make([][]int, M+1)
	var i, j int
	for i = 0; i <= M; i++ {
		dp[i] = make([]int, N+1)
	}
	for i = 1; i <= M; i++ {
		dp[i][0] = i
	}
	for j = 1; j <= N; j++ {
		dp[0][j] = j
	}
	for i = 0; i < M; i++ {
		for j = 0; j < N; j++ {
			if word1[i:i+1] == word2[j:j+1] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j], dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	return dp[M][N]
}
func main() {

}
