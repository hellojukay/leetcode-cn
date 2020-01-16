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

// https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	var first = head
	var last = head
	for {
		if last == nil {
			return false
		}
		for i := 0; i < 2; i++ {
			if first == nil {
				return false
			}
			first = first.Next
			if first == last {
				return true
			}
		}
		if last == first {
			return true
		}
		last = last.Next
	}
}
