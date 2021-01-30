package main

import "math"

const inf = math.MaxInt64

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func ladderLength(beginWord string, endWord string, wordList []string) (shortestPath int) {

	if beginWord == endWord {
		return 0
	}
	hash := map[string]int{}
	graph := [][]int{}

	addWord := func(word string) int {
		if _, ok := hash[word]; !ok {
			hash[word] = len(hash)
			graph = append(graph, []int{})
		}
		return hash[word]
	}

	addWordAndEdge := func(word string) {
		id := addWord(word)
		byteList := []byte(word)
		for i, b := range byteList {
			byteList[i] = '.'
			id2 := addWord(string(byteList))
			graph[id] = append(graph[id], id2)
			graph[id2] = append(graph[id2], id)
			byteList[i] = b
		}
	}

	for _, word := range wordList {
		addWordAndEdge(word)
	}
	addWordAndEdge(beginWord)

	if _, ok := hash[endWord]; !ok {
		return 0
	}

	dist := make([]int, len(hash))
	for i := range dist {
		dist[i] = inf
	}
	dist[hash[beginWord]] = 0

	q := []int{hash[beginWord]}
	var top int

	for len(q) > 0 {
		top, q = q[0], q[1:]
		if top == hash[endWord] {
			return dist[top]/2 + 1
		}
		for _, next := range graph[top] {
			if dist[next] == inf {
				dist[next] = dist[top] + 1
				q = append(q, next)
			}
		}
	}

	return 0
}

func main() {

}
