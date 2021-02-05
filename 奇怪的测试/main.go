package main

import (
	"fmt"
)

func test(data int) (ret int) {
	ret = data
	defer func() {
		ret++
	}()
	return ret
}

func main() {
	fmt.Println(test(512))
}
