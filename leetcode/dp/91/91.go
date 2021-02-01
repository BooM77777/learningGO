package main

import "strconv"

func numDecodings(s string) int {
	N := len(s)
	dp0, dp1 := 0, 1
	var tmp, res int
	var err error
	for i := 0; i < N; i++ {
		tmp = 0
		if res, err = strconv.Atoi(s[i : i+1]); err == nil && res > 0 {
			tmp += dp1
		}
		if i > 0 {
			if res, err = strconv.Atoi(s[i-1 : i+1]); err == nil && res >= 10 && res <= 26 {
				tmp += dp0
			}
		}
		if tmp == 0 {
			return 0
		}
		dp0, dp1 = dp1, tmp
	}
	return dp1
}
func main() {

}
