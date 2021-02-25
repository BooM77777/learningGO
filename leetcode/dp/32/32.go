package main

func longestValidParenthesesDP(s string) int {
	N := len(s)
	dp := make([]int, N)
	retVal := 0
	for i := 1; i < N; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if retVal < dp[i] {
				retVal = dp[i]
			}
		}
	}
	return retVal
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func longestValidParenthesesStack(s string) int {
	N := len(s)
	retVal, stack := 0, make([]int, 0, N+1)
	stack = append(stack, -1)
	for i := 0; i < N; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if stack[len(stack)-1] >= 0 && s[stack[len(stack)-1]] == '(' {
				retVal = max(retVal, i-stack[len(stack)-2])
			}
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			}
		}
		// fmt.Println(stack, retVal)
	}
	return retVal
}

func main() {

}
