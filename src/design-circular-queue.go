package main

type MyCircularQueue struct {
	head *queueNode
	tail *queueNode
	size int
	cap  int
}

type queueNode struct {
	next *queueNode
	pre  *queueNode
	val  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head: new(queueNode),
		tail: new(queueNode),
		cap:  k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		var newNode = &queueNode{
			val:  value,
			pre:  this.head,
			next: this.tail,
		}
		this.head.next = newNode
		this.tail.pre = newNode
		this.size++
		return true
	}
	var newNode = &queueNode{
		val:  value,
		pre:  this.tail.pre,
		next: this.tail,
	}
	this.tail.pre.next = newNode
	this.tail.pre = newNode
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.size == 1 {
		this.head.next = nil
		this.tail.pre = nil
		this.size--
		return true
	}
	this.head.next = this.head.next.next
	this.head.next.pre = this.head
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.next.val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.pre.val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
