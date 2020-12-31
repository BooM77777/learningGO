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
