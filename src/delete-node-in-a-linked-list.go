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

// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
