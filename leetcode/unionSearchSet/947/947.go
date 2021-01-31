package main

func removeStones(stones [][]int) int {
	N := len(stones)
	if N == 0 {
		return 0
	}
	root := make([]int, N)
	for i := range root {
		root[i] = i
	}
	var findRoot func(int) int
	findRoot = func(index int) int {
		if root[index] != index {
			r := findRoot(root[index])
			root[index] = r
			return r
		}
		return index
	}
	union := func(x, y int) {
		rX, rY := findRoot(x), findRoot(y)
		root[rX] = rY
	}
	hashX, hashY := map[int]int{}, map[int]int{}

	res, ok := -1, false
	for i, stone := range stones {
		if res, ok = hashX[stone[0]]; ok {
			union(i, res)
		} else {
			hashX[stone[0]] = i
		}
		if res, ok = hashY[stone[1]]; ok {
			union(i, res)
		} else {
			hashY[stone[1]] = i
		}
	}

	cnt := 0
	for i, r := range root {
		if r == i {
			cnt++
		}
	}
	return N - cnt
}

func main() {

}
