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

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	// 获取第二个节点，将他指向第一个节点，它将是新的头结点
	tmp := head.Next
	// 头节点指向第三节点后面的节点
	head.Next = swapPairs(head.Next.Next)
	tmp.Next = head
	// 第二个节点是头节点
	return tmp
}
