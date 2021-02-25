package main

import "fmt"

func fib() func() int {
	i, j := 0, 1
	return func() int {
		i, j = j, i+j
		return i
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
