package main

func merage(nums *[]int, left, mid, right int) int {
	retVal := 0
	tmp := make([]int, right-left+1)
	iter, iter1, iter2 := 0, left, mid+1
	for iter1 <= mid && iter2 <= right {
		if (*nums)[iter1] <= (*nums)[iter2] {
			tmp[iter] = (*nums)[iter1]
			iter, iter1 = iter+1, iter1+1
		} else {
			tmp[iter] = (*nums)[iter2]
			retVal += (mid + 1 - iter1)
			iter, iter2 = iter+1, iter2+1
		}
	}
	for iter1 <= mid {
		tmp[iter] = (*nums)[iter1]
		iter, iter1 = iter+1, iter1+1
	}
	for iter2 <= right {
		tmp[iter] = (*nums)[iter2]
		retVal += (mid + 1 - iter1)
		iter, iter2 = iter+1, iter2+1
	}
	copy((*nums)[left:right+1], tmp)
	return retVal
}

func merageSort(nums *[]int, left, right int) (retVal int) {
	if left < right {
		mid := (left + right) / 2
		retVal += merageSort(nums, left, mid) + merageSort(nums, mid+1, right)
		retVal += merage(nums, left, mid, right)
	}
	return
}

func reversePairs(nums []int) int {
	return merageSort(&nums, 0, len(nums)-1)
}

func main() {

}
