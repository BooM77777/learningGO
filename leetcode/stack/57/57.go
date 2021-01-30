package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {

	stack := make([][]int, 0, len(intervals)+1)

	myInsert := func(interval []int) {
		if len(stack) > 0 && stack[len(stack)-1][1] >= interval[0] {
			stack[len(stack)-1][1] = max(stack[len(stack)-1][1], interval[1])
		} else {
			stack = append(stack, interval)
		}
	}

	for _, interval := range intervals {
		if interval[0] > newInterval[0] {
			interval, newInterval = newInterval, interval
		}
		myInsert(interval)
	}
	myInsert(newInterval)
	return stack
}

func main() {

}
