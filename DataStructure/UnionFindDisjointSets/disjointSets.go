package main

import (
	"fmt"
)

//UnionFindDisjointSets 带权值的并查集
type UnionFindDisjointSets struct {
	size   int
	root   []int
	weight []float64
}

//NewUnionFindDisjointSets ...
func NewUnionFindDisjointSets(size int) *UnionFindDisjointSets {
	ret := &UnionFindDisjointSets{size: size}
	ret.root = make([]int, ret.size)
	ret.weight = make([]float64, ret.size)
	for i := 0; i < ret.size; i++ {
		ret.root[i] = -1
		ret.weight[i] = 1
	}
	return ret
}

//findRoot ...
func (s *UnionFindDisjointSets) findRoot(i int) (f int) {
	f = i
	if s.root[i] >= 0 {
		f = s.findRoot(s.root[i])
		s.weight[i] = s.weight[i] * s.weight[f]
		s.root[i] = f
	}
	return
}

//Union ...
func (s *UnionFindDisjointSets) Union(i, j int, res float64) {
	rooti := s.findRoot(i)
	rootj := s.findRoot(j)
	// 保证并查集树形结构的平衡，总是把子节点多的那个根节点作为新的根节点
	if s.root[rooti] < s.root[rootj] {
		i, j = j, i
		rooti, rootj = rootj, rooti
	}
	s.root[rooti], s.root[rootj] = s.root[i]+s.root[j], rooti
	s.weight[rooti] = s.weight[j] / s.weight[i] * res
}

//Find ...
func (s *UnionFindDisjointSets) Find(x, y int) (float64, bool) {
	if rootX, rootY := s.findRoot(x), s.findRoot(y); rootX == rootY {
		return s.weight[x] / s.weight[y], true
	}
	return -1.0, false
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	hash := map[string]int{}
	counter := 0
	for _, valPair := range equations {
		for _, val := range valPair {
			if _, ok := hash[val]; !ok {
				hash[val] = counter
				counter++
			}
		}
	}

	set := NewUnionFindDisjointSets(counter)
	for i, equation := range equations {
		v0, _ := hash[equation[0]]
		v1, _ := hash[equation[1]]
		set.Union(v0, v1, values[i])
	}

	ret := make([]float64, len(queries))
	for i, query := range queries {
		v0, ok0 := hash[query[0]]
		v1, ok1 := hash[query[1]]
		if ok0 && ok1 {
			ret[i], _ = set.Find(v0, v1)
		} else {
			ret[i] = -1
		}
	}

	return ret
}

func main() {
	// equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
	equations := [][]string{[]string{"a", "b"}, []string{"c", "d"}}
	values := []float64{1.0, 1.0}
	queries := [][]string{[]string{"a", "c"}, []string{"b", "d"}, []string{"b", "a"}, []string{"d", "c"}}
	fmt.Println(calcEquation(equations, values, queries))
}
