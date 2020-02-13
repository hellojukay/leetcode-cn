package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 */
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var tail *ListNode
	for {
		if head == nil {
			break
		}
		var tmp = head
		head = head.Next
		tmp.Next = tail
		tail = tmp
	}
	return tail
}
