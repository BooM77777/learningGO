package main

type (
	key interface{}
	val interface{}

	pair struct {
		k key
		v val
	}

	hashFunc func(k key) int

	//CuckooHash map
	CuckooHash struct {
		capSize      int
		bucket1      []*pair
		bucket2      []*pair
		hashFuncList []hashFunc
	}
)

//NewCuckooHash 构造函数
func NewCuckooHash(capSize int, f1, f2 hashFunc) *CuckooHash {
	ret := &CuckooHash{capSize, nil, nil, nil}
	ret.bucket1, ret.bucket2 = make([]*pair, capSize), make([]*pair, capSize)
	ret.hashFuncList = []hashFunc{f1, f2}
	return ret
}

//Insert 插入数据
func (h *CuckooHash) Insert(k key, v val) {

}

//Find 插入数据
func (h *CuckooHash) Find(k key) {

}

func main() {

}
