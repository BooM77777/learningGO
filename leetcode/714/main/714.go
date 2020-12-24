package main

import (
	"fmt"
)

// main 函数有问题
func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fmt.Println(maxProfit(prices, 2))
}

func maxProfit(prices []int, fee int) int {

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var empty = 0
	var busy = -prices[0]

	for _, price := range prices {
		empty = max(empty, busy+price-fee)
		busy = max(busy, empty-price)
	}

	return empty
}
