
type MyCircularDeque struct {
	head *node
	tail *node
	size int
	cap  int
}

type node struct {
	next *node
	pre  *node
	val  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		head: new(node),
		tail: new(node),
		cap:  k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		newNode := node{
			pre:  this.head,
			next: this.tail,
			val:  value,
		}
		this.head.next = &newNode
		this.tail.pre = &newNode
		this.size++
		return true
	}
	newNode := node{
		pre:  this.head,
		next: this.head.next,
		val:  value,
	}
	this.head.next.pre = &newNode
	this.head.next = &newNode
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		return this.InsertFront(value)
	}
	newNode := node{
		pre:  this.tail.pre,
		next: this.tail,
		val:  value,
	}
	this.tail.pre.next = &newNode
	this.tail.pre = &newNode
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.size == 1 {
		this.head.next = nil
		this.tail.pre = nil
		this.size = 0
		return true
	}
	this.size--
	this.head.next = this.head.next.next
	if this.head.next != nil {
		this.head.next.pre = this.head
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.size == 1 {
		return this.DeleteFront()
	}
	this.size--
	this.tail.pre = this.tail.pre.pre
	this.tail.pre.next = this.tail
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.next.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.pre.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.cap == this.size
}