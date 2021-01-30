package main

import (
	"sort"
)

// leetcode第235题

// BinaryIndexedTree 树状数组
type BinaryIndexedTree struct {
	a []int
	c []int
}

// NewBinaryIndexedTree 构造函数
func NewBinaryIndexedTree(size int) *BinaryIndexedTree {
	return &BinaryIndexedTree{
		a: make([]int, size+1),
		c: make([]int, size+1),
	}
}

// Query 查询操作
func (tree *BinaryIndexedTree) Query(id int) int {
	prefixSum := 0
	for id > 0 {
		prefixSum += tree.c[id]
		id -= id & (-id)
	}
	return prefixSum
}

// Insert 插入操作
func (tree *BinaryIndexedTree) Insert(id int) {
	tree.a[id]++
	for id < len(tree.c) {
		tree.c[id]++
		id += id & (-id)
	}
}

func countSmaller(nums []int) []int {

	// 去重
	hash := map[int]int{}
	for _, num := range nums {
		hash[num] = 1
	}
	// 排序
	i, bins := 0, make([]int, len(hash))
	for k := range hash {
		bins[i], i = k, i+1
	}
	sort.Ints(bins)
	// 构造与下标的映射关系
	for i, bin := range bins {
		hash[bin] = i + 1
	}

	res := make([]int, len(nums))
	tree := NewBinaryIndexedTree(len(bins))
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = tree.Query(hash[nums[i]])
		tree.Insert(hash[nums[i]])
	}

	return res
}

func main() {

}
