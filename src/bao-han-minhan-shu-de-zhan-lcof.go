type MinStack struct {
	data []Node
}

type Node struct {
	Val int
	Min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.data) == 0 {
		this.data = append(this.data, Node{Val: x, Min: x})
		return
	}
	min := this.data[len(this.data)-1].Min
	if x < min {
		min = x
	}
	this.data = append(this.data, Node{Val: x, Min: min})
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].Val
}

func (this *MinStack) Min() int {
	return this.data[len(this.data)-1].Min
}