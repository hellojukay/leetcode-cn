package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert(head *ListNode, n int) *ListNode {
	if head == nil {
		return &ListNode{
			Val: n,
		}
	}
	// match head with hd next
	if n < head.Val {
		node := &ListNode{
			Val:  n,
			Next: head,
		}
		return node
	}
	head.Next = insert(head.Next, n)
	return head
}
func insertionSortList(head *ListNode) *ListNode {
	var p *ListNode
	for {
		if head == nil {
			break
		}
		p = insert(p, head.Val)
		head = head.Next
	}
	return p
}
