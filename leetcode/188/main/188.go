package main

import (
	"fmt"
	"math"
)

func main() {
	k := 2
	prices := []int{2, 4, 7, 1}
	fmt.Println(maxProfit(k, prices))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a ...int) (ret int) {

	ret = a[0]
	for _, v := range a[1:] {
		if ret < v {
			ret = v
		}
	}
	return
}

func maxProfit(k int, prices []int) (ans int) {

	days := len(prices)

	k = min(k, days/2) // days 天最多可以完成 days / 2 笔交易

	dpBuy := make([][]int, days)
	dpSell := make([][]int, days)

	for i := 0; i < days; i++ {
		dpBuy[i] = make([]int, k+1)
		dpSell[i] = make([]int, k+1)
	}

	for i := 1; i <= k; i++ {
		dpBuy[0][i] = math.MinInt64 / 2
		dpSell[0][i] = math.MinInt64 / 2
	}

	dpBuy[0][0] = -prices[0]

	for i := 1; i < days; i++ {
		dpBuy[i][0] = max(dpBuy[i-1][0], dpSell[i-1][0]-prices[i])
		for j := 1; j < k+1; j++ {
			dpBuy[i][j] = max(dpSell[i-1][j]-prices[i], dpBuy[i-1][j])
			dpSell[i][j] = max(dpBuy[i-1][j-1]+prices[i], dpSell[i-1][j])
		}
	}

	ans = max(dpSell[days-1]...)
	return
}
