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

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var first = head
	var second = head
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first == nil {
		return second.Next
	}
	for {
		first = first.Next
		if first == nil {
			second.Next = second.Next.Next
			break
		}
		second = second.Next
	}
	return head
}
