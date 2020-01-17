package main

// https://leetcode-cn.com/problems/implement-stack-using-queues/
type MyStack struct {
	mycap int
	head  []int
	top   int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		mycap: 16,
		head:  make([]int, 16),
		top:   -1,
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.top == this.mycap-1 {
		// 扩容
		var newcap = this.mycap * 2
		var newHead = make([]int, newcap)
		for i := 0; i <= this.top; i++ {
			newHead[i] = this.head[i]
		}
		this.mycap = newcap
	}
	this.top++
	this.head[this.top] = x
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var result = this.head[this.top]
	this.top--
	return result
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.head[this.top]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.size() == 0
}
func (this *MyStack) size() int {
	return this.top + 1
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
