package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for j := 1; j <= target; j++ {
		for _, num := range nums {
			if j-num >= 0 {
				dp[j] += dp[j-num]
			}
		}
	}
	return dp[target]
}

func main() {

}
