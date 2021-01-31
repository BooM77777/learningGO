package main

import "math/rand"

type RandomizedSet struct {
	indexMap map[int]int
	set      []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		indexMap: map[int]int{},
		set:      []int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexMap[val]; !ok {
		this.indexMap[val], this.set = len(this.set), append(this.set, val)
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if tmpIndex, ok := this.indexMap[val]; ok {
		delete(this.indexMap, val)
		if tmpIndex != len(this.set)-1 {
			this.set[tmpIndex], this.set[len(this.set)-1] = this.set[len(this.set)-1], this.set[tmpIndex]
			this.indexMap[this.set[tmpIndex]] = tmpIndex
		}
		this.set = this.set[:len(this.set)-1]
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.set))
	return this.set[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {

}
