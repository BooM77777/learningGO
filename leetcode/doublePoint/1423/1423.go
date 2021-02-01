package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxScoreDP(cardPoints []int, k int) int {

	N := len(cardPoints)
	dp := make([][2]int, N+1)

	var i int

	for i = 0; i < N; i++ {
		dp[i+1][0] = dp[i][0] + cardPoints[i]
	}
	for i = N - 1; i >= 0; i-- {
		dp[N-i][1] = dp[N-i-1][1] + cardPoints[i]
	}

	ret := 0
	for i = 0; i <= k; i++ {
		left, right := i, k-i
		ret = max(ret, dp[left][0]+dp[right][1])
	}

	return ret
}

func maxScore(cardPoints []int, k int) int {

	var i int
	N := len(cardPoints)
	left, right, sum := 0, N-k, 0
	for i = left; i < right; i++ {
		sum += cardPoints[i]
	}
	ret := 0
	for i = right; i < len(cardPoints); i++ {
		ret += cardPoints[i]
	}
	total := ret + sum
	for left, right = left+1, right+1; right <= len(cardPoints); left, right = left+1, right+1 {
		sum += cardPoints[right-1] - cardPoints[left-1]
		ret = max(ret, total-sum)
	}

	return ret
}

func main() {

}
