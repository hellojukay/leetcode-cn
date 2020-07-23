package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var list1 []int
	var list2 []int
	for {
		if l1 == nil {
			break
		}
		list1 = append(list1, l1.Val)
		l1 = l1.Next
	}
	for {
		if l2 == nil {
			break
		}
		list2 = append(list2, l2.Val)
		l2 = l2.Next
	}
	var c = 0
	for {
		if len(list1) == 0 && len(list2) == 0 {
			break
		}
		if len(list1) > 0 {
			c = c + list1[len(list1)-1]
			list1 = list1[:len(list1)-1]
		}
		if len(list2) > 0 {
			c = c + list2[len(list2)-1]
			list2 = list2[:len(list2)-1]
		}
		var p = &ListNode{
			Val: c % 10,
		}
		c = c / 10
		if result == nil {
			result = p
			p.Next = nil
		} else {
			p.Next = result
			result = p
		}
	}
	if c > 0 {
		var p = &ListNode{
			Val: c,
		}
		p.Next = result
		result = p
	}
	return result
}
