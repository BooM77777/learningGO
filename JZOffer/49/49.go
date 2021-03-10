package main

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	iter2, iter3, iter5 := 0, 0, 0
	for i := 1; i < n; i++ {
		tmp2, tmp3, tmp5 := dp[iter2]*2, dp[iter3]*3, dp[iter5]*5
		if tmp2 < tmp3 {
			dp[i] = tmp2
		} else {
			dp[i] = tmp3
		}
		if dp[i] > tmp5 {
			dp[i] = tmp5
		}
		if dp[i] == tmp2 {
			iter2++
		}
		if dp[i] == tmp3 {
			iter3++
		}
		if dp[i] == tmp5 {
			iter5++
		}
	}
	return dp[len(dp)-1]
}

func main() {

}
