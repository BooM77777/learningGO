package main

import (
	"fmt"
	"sort"
)

func main() {
	list := [][]int{}
	fmt.Println(eraseOverlapIntervals(list))
}

func eraseOverlapIntervals(intervals [][]int) int {

	if len(intervals) <= 0 {
		return 0
	}
	// 使用结束时间排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	last := intervals[0]
	ans := 1
	for i := 1; i < len(intervals); i++ {
		if last[1] > intervals[i][0] {
			continue
		}
		ans++
		last = intervals[i]
	}
	return len(intervals) - ans
}

func eraseOverlapIntervalsDP(intervals [][]int) int {

	if len(intervals) <= 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	dp := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(intervals); i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return len(intervals) - max(dp...)
}

func max(a ...int) (ret int) {
	ret = a[0]
	for i := 1; i < len(a); i++ {
		if ret < a[i] {
			ret = a[i]
		}
	}
	return ret
}
