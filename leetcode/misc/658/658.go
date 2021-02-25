package main

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}

func findClosestElements(arr []int, k int, x int) []int {

	distArr := make([]int, len(arr))
	for i := range arr {
		distArr[i] = abs(arr[i] - x)
	}

	left, right := 0, len(arr)-1
	for right-left+1 > k {
		if abs(arr[left]-x) <= abs(arr[right]-x) {
			right--
		} else {
			left++
		}
	}

	return arr[left : right+1]
}

func main() {

}
