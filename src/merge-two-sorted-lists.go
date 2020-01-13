package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head = l1
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		head = l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		head = l2
	}
	return head
}
