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
	var result *ListNode
	for {
		if p == nil {
			break
		}
		if result == nil {
			result = p
			p = p.Next
			result.Next = nil
			continue
		} else {
			var tmp = p.Next
			p.Next = result
			result = p
			p = tmp
		}
	}
	return result
}
