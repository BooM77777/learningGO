package main

import (
	"fmt"
	"sync"
	"time"
)

func test(data int) (ret int) {
	ret = data
	defer func() {
		ret++
	}()
	return ret
}

func main() {
	var waitgroup sync.WaitGroup
	go func() {
		waitgroup.Add(1)
		fmt.Println("go 1")
		time.Sleep(1 * time.Second)
		waitgroup.Done()
	}()
	go func() {
		waitgroup.Add(1)
		fmt.Println("go 0")
		time.Sleep(1 * time.Second)
		waitgroup.Done()
	}()
	waitgroup.Wait()
	fmt.Println("Exit")
}
