package main

import (
	"fmt"
)

type distance struct {
	dst  int
	dist float64
}

func bfs(graph [][]float64, src, dst, counter int) float64 {
	q := []distance{}
	q = append(q, distance{src, 1})
	flag := make([]bool, counter)
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		flag[front.dst] = true
		if front.dst == dst {
			return front.dist
		}
		for i := 0; i < counter; i++ {
			if !flag[i] && graph[front.dst][i] != 0 {
				q = append(q, distance{i, front.dist * graph[front.dst][i]})
			}
		}
	}
	return -1.0
}

func calcEquationByFloyd(equations [][]string, values []float64, queries [][]string) []float64 {
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
	graph := make([][]float64, counter)
	for i := 0; i < counter; i++ {
		graph[i] = make([]float64, counter)
		for j := 0; j < counter; j++ {
			graph[i][j] = -1
		}
	}
	for i, valPair := range equations {
		val0, _ := hash[valPair[0]]
		val1, _ := hash[valPair[1]]
		graph[val0][val1] = values[i]
		graph[val1][val0] = 1 / values[i]
	}

	for i := 0; i < counter; i++ {
		for j := 0; j < counter; j++ {
			for k := 0; k < counter; k++ {
				if graph[i][j] == -1 && graph[i][k] != -1 && graph[k][j] != -1 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}
	ret := make([]float64, len(queries))

	for i, valPair := range queries {
		val0, ok0 := hash[valPair[0]]
		val1, ok1 := hash[valPair[1]]
		if ok0 && ok1 {
			ret[i] = graph[val0][val1]
		} else {
			ret[i] = -1
		}
	}
	return ret
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
	graph := make([][]float64, counter)
	for i := 0; i < counter; i++ {
		graph[i] = make([]float64, counter)
	}
	for i, valPair := range equations {
		val0, _ := hash[valPair[0]]
		val1, _ := hash[valPair[1]]
		graph[val0][val1] = values[i]
		graph[val1][val0] = 1 / values[i]
	}
	ret := make([]float64, len(queries))

	for i, valPair := range queries {
		val0, ok0 := hash[valPair[0]]
		val1, ok1 := hash[valPair[1]]
		if ok0 && ok1 {
			ret[i] = bfs(graph, val0, val1, counter)
		} else {
			ret[i] = -1.0
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
