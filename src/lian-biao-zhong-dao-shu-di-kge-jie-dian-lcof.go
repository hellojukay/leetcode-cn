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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var first, second = head, head
	for i := 0; i < k; i++ {
		first = first.Next
	}
	for {
		if first == nil {
			break
		}
		first = first.Next
		second = second.Next
	}
	return second
}
