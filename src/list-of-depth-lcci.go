package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/list-of-depth-lcci/
func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	var current []*TreeNode
	current = append(current, tree)
	var next []*TreeNode
	var result []*ListNode
	result = append(result, &ListNode{
		Val: tree.Val})
	for {
		if len(current) == 0 {
			break
		}
		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		if len(next) > 0 {
			var head = to_list(next)
			result = append(result, head)
		}
		current = next
		next = nil
	}
	return result
}
func to_list(arr []*TreeNode) *ListNode {
	var head, tail *ListNode
	for _, node := range arr {
		if head == nil {
			head = &ListNode{
				Val: node.Val,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: node.Val,
			}
			tail = tail.Next
		}
	}
	return head
}
