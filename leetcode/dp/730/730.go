package main

import "math"

func countPalindromicSubsequences(S string) int {
	N, base := len(S), int(math.Pow(10, 9))+7
	dp := [4][][]int{}
	sum := func(i, j int) int {
		return dp[0][i][j] + dp[1][i][j] + dp[2][i][j] + dp[3][i][j]
	}
	var i, j, k int
	for i = 0; i < 4; i++ {
		dp[i] = make([][]int, N)
		for j = 0; j < N; j++ {
			dp[i][j] = make([]int, N)
		}
	}
	for j = N - 1; j >= 0; j-- {
		for k = j; k < N; k++ {
			for i := 0; i < 4; i++ {
				chr := byte(i) + 'a'
				if j == k {
					if S[j] == chr {
						dp[i][j][k] = 1
					} else {
						dp[i][j][k] = 0
					}
				} else {
					if S[j] != chr {
						dp[i][j][k] = dp[i][j+1][k]
					} else if S[k] != chr {
						dp[i][j][k] = dp[i][j][k-1]
					} else {
						if j == k-1 {
							dp[i][j][k] = 2
						} else {
							dp[i][j][k] = (2 + sum(j+1, k-1)) % base
						}
					}
				}
			}
		}
	}

	return sum(0, N-1) % base
}
func main() {

}
