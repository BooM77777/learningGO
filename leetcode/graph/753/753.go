package main

import (
	"math"
	"strconv"
)

func crackSafe(n int, k int) string {
	seen, ans := map[int]bool{}, ""
	highest := int(math.Pow(10, float64(n-1)))

	var dfs func(int)
	dfs = func(node int) {
		var next int
		for i := 0; i < k; i++ {
			next = node*10 + i
			if !seen[next] {
				seen[next] = true
				dfs(next % highest)
				ans += strconv.Itoa(i)
			}
		}
	}
	dfs(0)

	for i := 0; i < n-1; i++ {
		ans = ans + "0"
	}
	return ans
}

func main() {

}
