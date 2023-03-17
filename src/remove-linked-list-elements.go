package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	var result *ListNode
	var tail *ListNode
	for {
		if head == nil {
			break
		}
		if head.Val != val {
			if result == nil {
				result = head
				head = head.Next
				tail = result
				tail.Next = nil
			} else {
				tail.Next = head
				head = head.Next
				tail = tail.Next
				tail.Next = nil
			}
		} else {
			head = head.Next
		}
	}
	return result
}

// 使用虚拟头结点
func removeElements2(head *ListNode, val int) *ListNode {
	var dummyHead = new(ListNode)
	dummyHead.Next = head
	var current = dummyHead
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummyHead.Next
}
