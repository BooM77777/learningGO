package main

func findContinuousSequence(target int) [][]int {

	totalNums := make([]int, target)
	for i := range totalNums {
		totalNums[i] = i + 1
	}

	retVal := [][]int{}
	l, r, win := 1, 1, 0
	for ; r < target && l <= target/2; r++ {
		win += r
		if win < target {
			continue
		}
		for win > target {
			win, l = win-l, l+1
		}
		if win == target {
			retVal = append(retVal, totalNums[l-1:r])
		}
	}

	return retVal
}

func main() {

}
