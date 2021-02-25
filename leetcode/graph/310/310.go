package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	counter := make([]int, n)
	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}
	for _, edge := range edges {
		counter[edge[0]]++
		counter[edge[1]]++
		graph[edge[0]][edge[1]] = true
		graph[edge[1]][edge[0]] = true
	}
	q := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if counter[i] == 1 {
			q = append(q, i)
		}
	}
	var res []int
	for len(q) > 0 {
		// fmt.Println(q)
		res = make([]int, 0, len(q))
		for _, item := range q {
			res = append(res, item)
			for i := 0; i < n; i++ {
				if graph[item][i] {
					counter[i]--
				}
			}
			counter[item] = n + 1
		}
		q = make([]int, 0, n)
		for i := 0; i < n; i++ {
			if counter[i] <= 1 {
				q = append(q, i)
			}
		}
	}
	return res
}

func main() {

}
