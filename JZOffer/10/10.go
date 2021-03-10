package main

func numWays(n int) int {
	base := 1000000007
	dp0, dp1 := 0, 1
	for i := 0; i < n; i++ {
		dp0, dp1 = dp1, dp0+dp1
		dp0, dp1 = dp0%base, dp1%base
	}
	return dp1
}

func main() {

}
