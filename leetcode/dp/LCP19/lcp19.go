package main

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func isYellow(color byte) int {
	if color == 'r' {
		return 0
	}
	return 1
}

func isRed(color byte) int {
	if color == 'r' {
		return 1
	}
	return 0
}

func minimumOperations(leaves string) int {
	const MAX = 10001
	dp0, dp1, dp2 := isYellow(leaves[0]), MAX, MAX
	for i := 1; i < len(leaves); i++ {
		if i >= 2 {
			dp0, dp1, dp2 = dp0+isYellow(leaves[i]), min(dp0, dp1)+isRed(leaves[i]), min(dp1, dp2)+isYellow(leaves[i])
		} else {
			dp0, dp1 = dp0+isYellow(leaves[i]), min(dp0, dp1)+isRed(leaves[i])
		}
	}
	return dp2
}

func main() {

}
