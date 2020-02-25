package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var tail *ListNode
	var head *ListNode
	for {
		if l1 == nil {
			tail.Next = l2
			break
		}
		if l2 == nil {
			tail.Next = l1
			break
		}
		var tmp *ListNode
		if l1.Val < l2.Val {
			tmp = l1
			l1 = l1.Next
		} else {
			tmp = l2
			l2 = l2.Next
		}
		if tail == nil {
			head = tmp
			tail = tmp
		} else {
			tail.Next = tmp
			tail = tail.Next
		}
	}
	tail.Next = nil
	return head
}
