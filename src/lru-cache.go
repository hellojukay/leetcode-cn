type LRUCache struct {
	capacity int
	size     int
	head     *LinkNode
	tail     *LinkNode
	cache    map[int]*LinkNode
}

type LinkNode struct {
	Val  int
	Key  int
	Pre  *LinkNode
	Next *LinkNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		head:     &LinkNode{},
		tail:     &LinkNode{},
		cache:    make(map[int]*LinkNode),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// remove
	if node != this.head.Next {
		node.Next.Pre = node.Pre
		node.Pre.Next = node.Next
		// insert into head
		this.head.Next.Pre = node
		node.Next = this.head.Next
		node.Pre = this.head
		this.head.Next = node
	}

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if this.size == 0 {
		node := &LinkNode{
			Val: value,
			Key: key,
		}
		this.size = 1
		this.head.Next = node
		node.Pre = this.head
		node.Next = this.tail
		this.tail.Pre = node
		this.cache[key] = node
		return
	}
	// key 已经存在的情况，删除数据，重新按照新增处理
	if node, ok := this.cache[key]; ok {
		delete(this.cache, node.Key)
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
		this.size--
	}
	node := &LinkNode{
		Key: key,
		Val: value,
	}
	this.head.Next.Pre = node
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next = node
	this.cache[key] = node
	this.size++
	if this.size > this.capacity {
		this.size--
		// remove from tail
		tail := this.tail.Pre
		delete(this.cache, tail.Key)
		this.tail.Pre = tail.Pre
		tail.Pre.Next = this.tail
	}
}