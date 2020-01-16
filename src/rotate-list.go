package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	if head == nil {
		return nil
	}
	var length = 1
	if head.Next == nil {
		return head
	}
	p := head
	// 链接首尾
	for {
		if p.Next == nil {
			p.Next = head
			break
		}
		p = p.Next
		length++
	}
	if k >= length {
		k = k % length
	}
	// 谁将成为新的头节点?
	// 倒数第 k 个节点会成为新的头节点
	for i := 0; i < length-k-1; i++ {
		head = head.Next
	}
	result := head.Next
	head.Next = nil
	return result
}

func main() {
	var node1 = new(ListNode)
	var node2 = new(ListNode)
	node1.Next = node2
	node1.Val = 1
	node2.Val = 2
	list := rotateRight(node1, 2)
	head := list
	for {
		if head == nil {
			break
		}
		println(head.Val)
		head = head.Next
	}
}
