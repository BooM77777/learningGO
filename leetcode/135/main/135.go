package main

import (
	"fmt"
)

func main() {
	ratings := []int{2, 0, 1}
	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {

	// 由于不考虑第一个人，因此初始的递增序列长度应该为1
	// 初始需要一个糖果，且第一个人分配到1个糖果
	inc, dec, ret, pre := 1, 0, 1, 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0 // 递减序列长度清0
			if ratings[i] == ratings[i-1] {
				// 重制
				pre = 1
			} else {
				// 递增
				pre = pre + 1
			}
			ret += pre
			inc = pre // 记录递增序列长度
		} else {
			// 进入递减序列
			dec++
			// 当递减序列长度恰好等于递增序列开始，
			// 就需要把递增序列尾端也计入递减序列，
			// 要不然递减序列就要-1了
			if dec == inc {
				dec++
			}
			ret += dec
			pre = 1
		}
	}
	return ret
}
