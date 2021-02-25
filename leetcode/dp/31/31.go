package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < target; i++ {
		for _, num := range nums {
			if i+num <= target {
				dp[i+num] += dp[i]
			}
		}
	}
	return dp[target]
}

func main() {

}
