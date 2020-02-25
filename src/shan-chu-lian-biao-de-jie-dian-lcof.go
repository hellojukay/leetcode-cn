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

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	var p = head
	var tail *ListNode
	for {
		if p == nil {
			break
		}
		if p.Val != val {
			if newHead == nil {
				newHead = p
				tail = newHead
			} else {
				tail.Next = p
				tail = tail.Next
			}
		}
		p = p.Next
	}
	tail.Next = nil
	return newHead
}
