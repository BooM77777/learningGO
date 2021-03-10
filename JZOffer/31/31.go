package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	iter1, iter2 := 0, 0
	for iter2 < len(popped) {
		if len(stack) > 0 && stack[len(stack)-1] == popped[iter2] {
			stack, iter2 = stack[:len(stack)-1], iter2+1
		} else {
			if iter1 == len(pushed) {
				return false
			}
			stack, iter1 = append(stack, pushed[iter1]), iter1+1
		}
	}
	return true
}

func main() {

}
