package main

func minSwaps(data []int) int {
	N := len(data)
	winSize := 0
	for i := range data {
		if data[i] == 1 {
			winSize++
		}
	}
	win := 0
	for i := 0; i < winSize; i++ {
		if data[i] == 0 {
			win++
		}
	}
	ret := win
	for i := winSize; i < N; i++ {
		if data[i-winSize] == 0 {
			win--
		}
		if data[i] == 0 {
			win++
		}
		if win < ret {
			ret = win
		}
	}
	return ret
}

func main() {

}
