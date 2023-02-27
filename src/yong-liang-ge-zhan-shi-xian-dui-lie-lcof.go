package main

type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 && len(this.inStack) == 0 {
		return -1
	}
	if len(this.outStack) == 0 {
		for j := len(this.inStack) - 1; j >= 0; j-- {
			this.outStack = append(this.outStack, this.inStack[j])
		}
		this.inStack = nil
	}
	last := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return last
}
