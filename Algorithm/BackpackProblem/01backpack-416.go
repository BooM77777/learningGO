func canPartition(nums []int) bool {
	sum, maxVal := 0, nums[0]
	for _, num := range nums {
		sum += num
		if maxVal < num {
			maxVal = num
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if target < maxVal {
		return false
	}
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j > 0; j-- {
			if j-num >= 0 {
				dp[j] = dp[j] || dp[j-num]
			}
		}
	}
	return dp[target]
}

func main() {

}