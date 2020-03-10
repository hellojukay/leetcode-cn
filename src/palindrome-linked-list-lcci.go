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

// https://leetcode-cn.com/problems/palindrome-linked-list-lcci/
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
		if head.Val == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, head.Val)
		}
		head = head.Next
	}
	return len(stack) == 0
}
