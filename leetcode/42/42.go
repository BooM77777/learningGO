package main

import (
	"fmt"
)

func min(input ...int) (ret int) {
	ret = input[0]
	for i := 1; i < len(input); i++ {
		if ret > input[i] {
			ret = input[i]
		}
	}
	return
}

func trapStack(height []int) (ret int) {
	stack := []int{}
	curr := 0
	for curr < len(height) {
		for len(stack) > 0 && height[curr] > height[stack[len(stack)-1]] {
			base := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) <= 0 {
				break
			}
			height := min(height[curr], height[stack[len(stack)-1]]) - base
			width := curr - stack[len(stack)-1] - 1
			ret = ret + height*width
		}
		stack = append(stack, curr)
		curr++
	}
	return
}

func main() {
	l := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trapStack(l))
}
