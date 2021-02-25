package main

func minimumSwap(s1 string, s2 string) int {
	N := len(s1)
	xyCnt, yxCnt := 0, 0
	for i := 0; i < N; i++ {
		if s1[i] == 'x' && s2[i] == 'y' {
			xyCnt++
		} else if s1[i] == 'y' && s2[i] == 'x' {
			yxCnt++
		}
	}
	if (xyCnt+yxCnt)%2 == 1 {
		return -1
	}
	retVal := xyCnt/2 + yxCnt/2
	if xyCnt%2 == 1 && yxCnt%2 == 1 {
		retVal += 2
	}
	return retVal
}

func main() {

}
