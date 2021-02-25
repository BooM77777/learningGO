package main

import "fmt"

// MergeSort ...
func MergeSort(nums *[]int) {
	mergeSort(nums, 0, len(*nums))
}

func mergeSort(nums *[]int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(nums, left, mid)
		mergeSort(nums, mid, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums *[]int, left, mid, right int) {
	tmp := make([]int, right-left)
	iter, iter0, iter1 := 0, left, mid
	for iter0 < mid && iter1 < right {
		if (*nums)[iter0] <= (*nums)[iter1] {
			tmp[iter] = (*nums)[iter0]
			iter, iter0 = iter+1, iter0+1
		} else {
			tmp[iter] = (*nums)[iter1]
			iter, iter1 = iter+1, iter1+1
		}
	}
	for iter0 < mid {
		tmp[iter] = (*nums)[iter0]
		iter, iter0 = iter+1, iter0+1
	}
	for iter1 < right {
		tmp[iter] = (*nums)[iter1]
		iter, iter1 = iter+1, iter1+1
	}
	// fmt.Println(left, mid, right)
	copy((*nums)[left:right], tmp)
}

// func mergeSort(nums *[]int, left, right int) {
// 	if left < right {
// 		mid := (left + right) / 2
// 		mergeSort(nums, left, mid)
// 		mergeSort(nums, mid+1, right)
// 		merge(nums, left, mid, right)
// 	}
// }

// func merge(nums *[]int, left, mid, right int) {
// 	tmp := make([]int, right-left+1)
// 	iter, iter1, iter2 := 0, left, mid+1
// 	for iter1 <= mid && iter2 <= right {
// 		if (*nums)[iter1] <= (*nums)[iter2] {
// 			tmp[iter] = (*nums)[iter1]
// 			iter, iter1 = iter+1, iter1+1
// 		} else {
// 			tmp[iter] = (*nums)[iter2]
// 			iter, iter2 = iter+1, iter2+1
// 		}
// 	}
// 	for iter1 <= mid {
// 		tmp[iter] = (*nums)[iter1]
// 		iter, iter1 = iter+1, iter1+1
// 	}
// 	for iter2 <= right {
// 		tmp[iter] = (*nums)[iter2]
// 		iter, iter2 = iter+1, iter2+1
// 	}
// 	copy((*nums)[left:right+1], tmp)
// }

func main() {
	// nums := []int{2, 5, 8, 4, 3, 7, 6, 9, 10, 11, 2, 5}
	nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	// left, right := 0, len(nums)
	// mid := (right + left) / 2
	// merge(&nums, left, mid, right)
	MergeSort(&nums)
	fmt.Println(nums)
}
