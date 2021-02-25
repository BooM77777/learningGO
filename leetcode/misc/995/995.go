package main

func minKBitFlips(A []int, K int) int {
	reversCnt, ans := 0, 0
	for i := range A {
		if i-K >= 0 && A[i-K] >= 2 {
			reversCnt--
		}
		if (A[i]+reversCnt)&1 == 0 {
			if i+K <= len(A) {
				A[i] += 2
				reversCnt, ans = reversCnt+1, ans+1
			} else {
				return -1
			}
		}
	}
	return ans
}

func main() {

}
