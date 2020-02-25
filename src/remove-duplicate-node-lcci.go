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

// https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var tail = head
	var p = head
	var has = make(map[int]bool)
	has[tail.Val] = true
	for {
		if p == nil {
			break
		}
		if has[p.Val] {
			p = p.Next
		} else {
			tail.Next = p
			tail = tail.Next
			has[p.Val] = true
			p = p.Next
			tail.Next = nil
		}
	}
	return head
}
