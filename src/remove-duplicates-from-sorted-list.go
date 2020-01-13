package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var p = head
	var q = p.Next
	for {
		p.Next = nil
		// find first different value
		for {
			// link end
			if q == nil {
				break
			}
			if q.Val != p.Val {
				break
			}
			q = q.Next
		}
		// link end
		if q == nil {
			break
		}
		if p.Val != q.Val {
			p.Next = q
			p = p.Next
			q = p.Next
		}
	}
	return head
}
