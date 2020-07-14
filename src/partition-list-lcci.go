package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/partition-list-lcci/
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var lt []*ListNode
	var gt []*ListNode
	for {
		if head == nil {
			break
		}
		if head.Val < x {
			lt = append(lt, head)
		} else {
			gt = append(gt, head)
		}
		head = head.Next
	}
	var list = append(lt, gt...)
	var p *ListNode
	head = list[0]
	p = head
	for i, node := range list {
		if i == 0 {
			continue
		}
		p.Next = node
		p = p.Next
	}
	p.Next = nil
	return head
}
