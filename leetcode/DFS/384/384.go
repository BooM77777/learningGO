package main

func lexicalOrder(n int) []int {
	max := n
	data := make([]int, 0, n)
	var preOrder func(n int)
	preOrder = func(n int) {
		if n == 0 || n > max {
			return
		}
		data = append(data, n)
		var tmpNum int
		for i := 0; i < 10; i++ {
			tmpNum = n*10 + i
			if tmpNum <= max {
				preOrder(tmpNum)
			}
		}
	}
	for i := 1; i < 10; i++ {
		preOrder(i)
	}
	return data
}
func main() {

}
