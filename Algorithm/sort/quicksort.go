package main

import "fmt"

// QuickSort ...
func QuickSort(nums *[]int) {
	quickSort(nums, 0, len(*nums)-1)
}

func quickSort(nums *[]int, left, right int) {
	if left < right {
		mid := getPosition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

func getPosition(nums *[]int, left, right int) int {
	midVal := (*nums)[left]
	for left < right {
		for left < right && (*nums)[right] >= midVal {
			right--
		}
		if left < right {
			(*nums)[left] = (*nums)[right]
			left++
		}
		for left < right && (*nums)[left] <= midVal {
			left++
		}
		if left < right {
			(*nums)[right] = (*nums)[left]
			right--
		}
	}
	(*nums)[left] = midVal
	return left
}

func main() {
	nums := []int{2, 5, 8, 4, 3, 7, 6, 9, 10, 11, 2, 5}
	QuickSort(&nums)
	fmt.Println(nums)
}
