/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode-cn.com/problems/add-two-numbers/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var sum = l1.Val + l2.Val
	var newNode = &ListNode{
		Val:  sum % 10,
		Next: nil,
	}
	if sum >= 10 {
		newNode.Next = addTwoNumbers(addTwoNumbers(&ListNode{Val: 1}, l1.Next), l2.Next)
	} else {
		newNode.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return newNode
}
