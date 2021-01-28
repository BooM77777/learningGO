package main

import "fmt"

func reserve(nums []int, begin, end int) {
	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// for some reson
func nextPermutation(nums []int) {
	N := len(nums)
	var iter0, iter1 int
	for iter0 = N - 2; iter0 >= 0; iter0-- {
		if nums[iter0] < nums[iter0+1] {
			break
		}
	}
	if iter0 < 0 {
		reserve(nums, 0, N-1)
		return
	}
	for iter1 = N - 1; iter1 >= 0; iter1-- {
		if nums[iter1] > nums[iter0] {
			break
		}
	}
	nums[iter0], nums[iter1] = nums[iter1], nums[iter0]
	reserve(nums, iter0+1, N-1)
}

func main() {
	nums := []int{4, 5, 2, 6, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
