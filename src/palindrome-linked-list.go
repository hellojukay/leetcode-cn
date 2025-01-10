package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	var stack []int
	for {
		if head == nil {
			break
		}
		if len(stack) == 0 {
			stack = append(stack, head.Val)
			head = head.Next
			continue
		}
		top := stack[len(stack)-1]
		if head.Val == top {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, head.Val)
		}
		head = head.Next
	}
	return r.String() == reverse(r.String())
}
