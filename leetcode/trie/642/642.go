package main

type (
	Sentence struct {
		str  string
		heat int
	}
	TrieNode struct {
		next [128]*TrieNode
		sent *Sentence
	}

	AutocompleteSystem struct {
		root   *TrieNode
		prefix []byte
	}
)

func NewSentence(sentence string, heat int) Sentence {
	return Sentence{
		str:  sentence,
		heat: heat,
	}
}

func (s *Sentence) less(s1 *Sentence) bool {
	if s.heat == s1.heat {
		for i, j := 0, 0; i < len(s.str) && j < len(s1.str); i, j = i+1, j+1 {
			if s.str[i] == s1.str[j] {
				continue
			}
			return s.str[i] > s1.str[j]
		}
		return len(s.str) > len(s1.str)
	}
	return s.heat < s1.heat
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func (as *AutocompleteSystem) insert(sentence Sentence) {
	tmpNode := as.root
	for _, b := range sentence.str {
		if tmpNode.next[int(b)] == nil {
			tmpNode.next[int(b)] = NewTrieNode()
		}
		tmpNode = tmpNode.next[int(b)]
	}
	if tmpNode.sent == nil {
		tmpNode.sent = &sentence
	} else {
		tmpNode.sent.heat++
	}
}

func (as *AutocompleteSystem) search() []*Sentence {

	ret, tmpNode := make([]*Sentence, 0, 100), as.root
	for _, b := range as.prefix {
		if tmpNode.next[int(b)] == nil {
			return ret
		}
		tmpNode = tmpNode.next[int(b)]
	}

	// BFS
	q, top := []*TrieNode{tmpNode}, tmpNode
	for len(q) > 0 {
		top, q = q[0], q[1:]
		if top.sent != nil {
			ret = append(ret, top.sent)
		}
		for _, next := range top.next {
			if next != nil {
				q = append(q, next)
			}
		}
	}
	return ret
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	if len(sentences) != len(times) {
		return AutocompleteSystem{}
	}
	N := len(sentences)
	as := AutocompleteSystem{
		root:   NewTrieNode(),
		prefix: make([]byte, 0, 100),
	}
	for i := 0; i < N; i++ {
		as.insert(NewSentence(sentences[i], times[i]))
	}
	return as
}

func bubbleSort(sentList []*Sentence) string {

	for i := 1; i < len(sentList); i++ {
		if sentList[i].less(sentList[i-1]) {
			sentList[i], sentList[i-1] = sentList[i-1], sentList[i]
		}
	}
	return sentList[len(sentList)-1].str
}

func (this *AutocompleteSystem) Input(c byte) []string {
	ret := make([]string, 0, 3)
	if c == '#' {
		this.insert(NewSentence(string(this.prefix), 1))
		this.prefix = make([]byte, 0, 100)
	} else {
		this.prefix = append(this.prefix, c)
		sentList := this.search()
		// for _, sent := range sentList {
		//     fmt.Println(sent)
		// }
		// 3次冒泡排序
		for i := 0; i < 3; i++ {
			if len(sentList) == 0 {
				break
			}
			ret, sentList = append(ret, bubbleSort(sentList)), sentList[:len(sentList)-1]
		}
		// fmt.Println(ret)
		// fmt.Println("====")
	}
	return ret
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

func main() {

}
