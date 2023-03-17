type MyLinkedList struct {
	Next *MyLinkedList
	Val  int
}

func Constructor() MyLinkedList {
	var dummyHead = MyLinkedList{}
	return dummyHead
}

func (this *MyLinkedList) Get(index int) int {
	var i = 0
	var current = this.Next
	for current != nil {
		if i == index {
			return current.Val
		}
		current = current.Next
		i++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := new(MyLinkedList)
	node.Val = val
	node.Next = this.Next
	this.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	var current = this
	for current != nil {
		if current.Next == nil {
			current.Next = &MyLinkedList{
				Val: val,
			}
			break
		}
		current = current.Next
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	var i = 0
	var current = this
	for current != nil {
		if i == index {
			node := new(MyLinkedList)
			node.Val = val
			node.Next = current.Next
			current.Next = node
			break
		}
		i++
		current = current.Next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	var i = 0
	var current = this
	for current != nil {
		if i == index {
			if current.Next != nil {
				current.Next = current.Next.Next
				return
			}
		}
		i++
		current = current.Next
	}
}