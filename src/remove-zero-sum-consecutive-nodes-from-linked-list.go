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

// https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var stack = make([]*ListNode, 1001)
	var top = -1
	for {
		if head == nil {
			break
		}
		// 跳过 0 值节点
		if head.Val == 0 {
			head = head.Next
			continue
		}
		// checksum
		index := find_sum(stack, top, -1*head.Val)
		if index != -1 {
			// 去掉可以被弹出的部分
			top = index - 1
			head = head.Next
			continue
		}
		if top == -1 {
			top++
			stack[top] = head
			head = head.Next
			continue
		}
		stack[top].Next = head
		top++
		stack[top] = head
		head = head.Next
	}
	if top >= 0 {
		stack[top].Next = nil
	} else {
		return nil
	}
	return stack[0]
}

// 检查当且栈是否可以弹出，如果可以则返回可以弹出的高度
// 否则返回 -1
func find_sum(stack []*ListNode, top int, value int) int {
	if top == -1 {
		return -1
	}
	index := -1
	var sum = 0
	for i := top; i >= 0; i-- {
		sum = sum + stack[i].Val
		if sum == value {
			index = i
			break
		}
	}
	return index
}
