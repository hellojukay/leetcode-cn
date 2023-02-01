package main

type element struct {
	data        int
	currentMint int
}
type MinStack struct {
	stack []element
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, element{data: val, currentMint: val})
		return
	}
	var min = this.stack[len(this.stack)-1].currentMint
	if val < min {
		min = val
	}
	this.stack = append(this.stack, element{data: val, currentMint: min})

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].data
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].currentMint

}
