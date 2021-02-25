package main

import "math"

func increasingTriplet(nums []int) bool {
	small, mid := math.MaxInt64, math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] <= small {
			small = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}

func main() {

}
