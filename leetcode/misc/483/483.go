package main

func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	num := make([]int, 2*n)
	num[0] = 1
	retVal, j := 0, 1
	for i := 1; i < n && j < n; i++ {
		if num[i] == 1 {
			num[j] = num[j-1] ^ 3
			j++
		} else {
			num[j] = num[j-1] ^ 3
			num[j+1] = num[j-1] ^ 3
			j += 2
		}
	}
	for i := 0; i < n; i++ {
		if num[i] == 1 {
			retVal += 1
		}
	}
	return retVal
}

func main() {

}
