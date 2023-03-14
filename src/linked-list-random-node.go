/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	cache map[int]*ListNode
	size int 
}

func Constructor(head *ListNode) Solution {
	var s =  Solution{
		cache: make(map[int]*ListNode)
	}
	var i int 
	for head != nil {
		cache[i] = head
		i++
		head = head.Next
		s.size++
	}
}

func (this *Solution) GetRandom() int {
	var n = rand.Int(this.size)
	node := cache[n]
	return node.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */