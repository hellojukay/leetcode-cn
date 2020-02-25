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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var m = make(map[*ListNode]bool)
	for {
		if headA == nil && headB == nil {
			break
		}
		if headA != nil {
			if m[headA] {
				return headA
			}
			m[headA] = true
			headA = headA.Next
		}
		if headB != nil {
			if m[headB] {
				return headB
			}
			m[headB] = true
			headB = headB.Next
		}
	}
	return nil
}
