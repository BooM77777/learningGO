package main

func climbStairs(n int) int {
	dp0, dp1 := 0, 1
	for i := 0; i < n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp1
}

func main() {

}
