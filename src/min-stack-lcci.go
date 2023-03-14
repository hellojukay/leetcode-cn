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
		node := Node{
			Val: x,
			Min: x,
		}
		this.data = append(this.data, node)
		return
	}
	min := this.data[len(this.data)-1].Min
	if x < min {
		min = x
	}
	node := Node{
		Val: x,
		Min: min,
	}
	this.data = append(this.data, node)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1].Min

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */