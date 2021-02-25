package main

func isBipartite(graph [][]int) bool {

	colorMap := make([]int, len(graph))

	var dfs func(node int, color int) bool
	dfs = func(node int, color int) bool {
		colorMap[node] = color
		for _, nextNode := range graph[node] {
			if colorMap[nextNode] == color {
				return false
			} else if colorMap[nextNode] == 0 {
				if dfs(nextNode, color^3) == false {
					return false
				}
			}
		}
		return true
	}

	for i := range colorMap {
		if colorMap[i] == 0 {
			if dfs(i, 1) == false {
				return false
			}
			// fmt.Println(colorMap)
		}
	}
	return true
}

func main() {

}
