package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) (ret int) {
	N := len(heights)
	var stack []int
	left, right := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		left[i], right[i] = 0, N
	}
	stack = make([]int, 0, N)
	for i := 0; i < N; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1] + 1
		}
		stack = append(stack, i)
	}
	for i := 0; i < N; i++ {
		ret = max(ret, (right[i]-left[i])*heights[i])
	}
	return ret
}

func main() {

}
