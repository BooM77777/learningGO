package main

func isMatch(s string, p string) bool {
	match := func(i, j int) bool {
		if i == 0 {
			return false
		} else if p[j-1] == '.' {
			return true
		} else {
			return s[i-1] == p[j-1]
		}
	}
	M, N := len(s), len(p)
	dp := make([][]bool, M+1)
	for i := range dp {
		dp[i] = make([]bool, N+1)
	}
	dp[0][0] = true
	for j := 1; j <= N; j++ {
		for i := 0; i <= M; i++ {
			if p[j-1] == '*' {
				// 不匹配的情况，跳过*与前一个字符所以是j-2
				dp[i][j] = dp[i][j] || dp[i][j-2]
				// 匹配一个的情况就是匹配当前字符和*的前一个字符
				if match(i, j-1) {
					// 如果可以匹配，就把当前字符扔掉，匹配上一个
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if match(i, j) {
				// 这个应该懂
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[M][N]
}

func main() {

}
