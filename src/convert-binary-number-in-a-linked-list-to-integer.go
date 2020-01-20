package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/
func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	var result = 0
	var stack = make([]int, 30)
	var length = 0
	for {
		if head == nil {
			break
		}
		length++
		stack = append(stack, head.Val)
		head = head.Next
	}
	n := 0
	for {
		if len(stack) == 0 {
			break
		}
		result = result + pow2(n)*stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n++
	}
	return result
}
func pow2(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result = result * 2
	}
	return result
}
