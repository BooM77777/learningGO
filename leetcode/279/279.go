package main

import "fmt"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 5
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			tmp := j * j
			if tmp > i {
				break
			}
			if dp[i-tmp]+1 < dp[i] {
				dp[i] = dp[i-tmp] + 1
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(998))
}
