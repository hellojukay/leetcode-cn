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

func detectCycle(head *ListNode) *ListNode {
	var first = head
	var m = make(map[*ListNode]bool)
	for {
		if first == nil {
			return nil
		}
		if m[first] {
			return first
		}
		m[first] = true
		first = first.Next
	}
}
