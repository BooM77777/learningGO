package main

import "math"

func kEmptySlots(bulbs []int, k int) int {
	if k > len(bulbs)-2 {
		return -1
	}
	bulbList := make([]int, len(bulbs))
	for i, bulb := range bulbs {
		bulbList[bulb-1] = i + 1
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	retVal := math.MaxInt64
	left, right := 0, k+1
search:
	for right < len(bulbList) {
		for i := left + 1; i < right; i++ {
			if bulbList[i] < bulbList[left] || bulbList[i] < bulbList[right] {
				left, right = i, i+k+1
				continue search
			}
		}
		retVal = min(retVal, max(bulbList[left], bulbList[right]))
		left, right = right, right+k+1
	}
	if retVal == math.MaxInt64 {
		return -1
	}
	return retVal
}

func main() {

}
