package main

// CQueue ...
type CQueue struct {
	stack1 []int
	stack2 []int
}

// Constructor ...
func Constructor() CQueue {
	return CQueue{
		stack1: []int{},
		stack2: []int{},
	}
}

// AppendTail ...
func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

// DeleteHead ...
func (this *CQueue) DeleteHead() (retVal int) {
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	if len(this.stack2) == 0 {
		return -1
	}
	retVal = this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return
}

func main() {

}
