package main

func isScramble(s1 string, s2 string) bool {

	N, M := len(s1), len(s2)
	if N != M {
		return false
	}
	// dp
	dp := make([][][]bool, N)
	for i := range dp {
		dp[i] = make([][]bool, N)
		for j := range dp[i] {
			dp[i][j] = make([]bool, N+1)
			dp[i][j][0] = true
			dp[i][j][1] = (s1[i] == s2[j])
		}
	}

	for l := 2; l <= N; l++ {
		for i := 0; i <= N-l; i++ {
			for j := 0; j <= N-l; j++ {
				for k := 1; k < l; k++ {

					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}

					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][N]
}

func main() {

}
