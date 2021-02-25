package main

func longestOnes(A []int, K int) int {
	left, right, zeroCnt, ret := 0, 0, 0, 0
	for ; right < len(A); right++ {
		if A[right] == 0 {
			zeroCnt++
		}
		for left <= right && zeroCnt > K {
			if A[left] == 0 {
				zeroCnt--
			}
			left++
		}
		if right-left+1 > ret {
			ret = right - left + 1
		}
	}
	return ret
}

func main() {

}
