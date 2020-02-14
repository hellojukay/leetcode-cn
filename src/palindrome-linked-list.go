package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	var r = reverse(head)
	for {
		if head == nil {
			break
		}
		if r.Val != head.Val {
			return false
		}
		r = r.Next
		head = head.Next
	}
	return true
}
func reverse(root *ListNode) *ListNode {
	if root == nil {
		return nil
	}
	var head *ListNode
	for {
		if root == nil {
			break
		}
		if head == nil {
			head = &ListNode{
				Val: root.Val,
			}
			root = root.Next
			continue
		}
		var tmp = &ListNode{
			Val: root.Val,
		}
		tmp.Next = head
		head = tmp
		root = root.Next
	}
	return head
}
