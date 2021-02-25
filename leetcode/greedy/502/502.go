package main

import "math"

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	n := len(Profits)
	iter := 0
	for ; iter < k; iter++ {
		index := -1
		for j := 0; j < n; j++ {
			if W >= Capital[j] {
				if index == -1 || Profits[index] < Profits[j] {
					index = j
				}
			}
		}
		if index == -1 {
			break
		}
		W += Profits[index]
		Capital[index] = math.MaxInt64
	}
	return W
}

func main() {

}
