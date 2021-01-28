package main

import "fmt"

type (
	Trie struct {
		root *TrieNode
	}
	TrieNode struct {
		next [26]*TrieNode
		word string
	}
)

func NewTrie() (ret *Trie) {
	ret = &Trie{root: &TrieNode{word: ""}}
	return
}

func (t *Trie) Insert(word string) {
	iter := t.root
	for i := 0; i < len(word); i++ {
		if iter.next[int(word[i]-'a')] == nil {
			iter.next[int(word[i]-'a')] = &TrieNode{word: ""}
		}
		iter = iter.next[int(word[i]-'a')]
	}
	iter.word = word
}

func (t *Trie) SearchByPrefix(prefix string) (ret []string) {
	tmp := t.root
	var next *TrieNode
	for i := 0; i < len(prefix); i++ {
		next = tmp.next[int(prefix[i]-'a')]
		if next == nil {
			return
		}
		tmp = next
	}
	stack := []*TrieNode{tmp}
	for len(stack) > 0 {
		tmp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmp.word != "" {
			ret = append(ret, tmp.word)
		}
		for i := 0; i < 26; i++ {
			if tmp.next[i] != nil {
				stack = append(stack, tmp.next[i])
			}
		}
	}
	return
}

func function(words []string, N int, n int, prefix string, trie *Trie, ret *[][]string) {
	if N == n {
		tmp := make([]string, len(words))
		copy(tmp, words)
		(*ret) = append((*ret), tmp)
		return
	}
	wordList := trie.SearchByPrefix(prefix)
	flag := true
	for _, word := range wordList {
		for i := 0; i < n; i++ {
			if word == words[i] {
				flag = false
			}
		}
		if flag {
			words[n] = word
			n++
			prefix = ""
			for i := 0; n < N && i < n; i++ {
				prefix += words[i][n : n+1]
			}
			function(words, N, n, prefix, trie, ret)
			n--
			for i := 0; i < n; i++ {
				prefix += words[i][n : n+1]
			}
		}
	}
}

func wordSquares(words []string) (ret [][]string) {
	trie := NewTrie()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	ret = [][]string{}
	N := len(words[0])
	function(make([]string, N), N, 0, "", trie, &ret)
	return ret
}

func main() {
	words := []string{"area", "lead", "wall", "lady", "ball"}
	fmt.Println(wordSquares(words))
}
