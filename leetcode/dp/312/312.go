package main

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func maxCoins(nums []int) int {
	N := len(nums)
	dp := make([][]int, N+2)
	for i := 0; i < N+2; i++ {
		dp[i] = make([]int, N+2)
	}
	val := make([]int, N+2)
	val[0], val[N+1] = 1, 1
	for i := 0; i < N; i++ {
		val[i+1] = nums[i]
	}
	for i := N - 1; i >= 0; i-- {
		for j := i + 2; j < N+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+val[i]*val[k]*val[j]+dp[k][j])
			}
		}
	}
	return dp[0][N+1]
}
func main() {

}
