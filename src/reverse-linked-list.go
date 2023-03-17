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

// https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p = head
	var newHead *ListNode
	for p != nil {
		var next = p.Next
		p.Next = newHead
		newHead = p
		p = next
	}
	return newHead
}
