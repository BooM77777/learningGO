package main

func main() {

}

func minCostClimbingStairs(cost []int) int {
	var min func(a, b int) int
	min = func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	var n = len(cost)
	var prev = 0
	var curr = 0
	for i := 2; i <= n; i++ {
		curr, prev = min(prev+cost[i-2], curr+cost[i-1]), curr
	}
	return curr
}
