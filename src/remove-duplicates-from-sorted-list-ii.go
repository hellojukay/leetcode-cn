package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var stack []*ListNode
	var needDelete = false
	for {
		if head == nil {
			if needDelete {
				if len(stack) == 1 {
					stack = nil
				} else {
					stack = stack[:len(stack)-1]
				}
			}
			break
		}
		if len(stack) == 0 {
			stack = append(stack, head)
			head = head.Next
			continue
		}
		if stack[len(stack)-1].Val == head.Val {
			needDelete = true
			head = head.Next
			continue
		}
		if stack[len(stack)-1].Val != head.Val && needDelete {
			if len(stack) == 1 {
				stack = nil
				stack = append(stack, head)
				needDelete = false
				head = head.Next
				continue
			}
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				stack[len(stack)-1].Next = head
			}
			stack = append(stack, head)
			head = head.Next
			needDelete = false
			continue
		}
		stack[len(stack)-1].Next = head
		stack = append(stack, head)
		head = head.Next
	}
	if len(stack) == 0 {
		return nil
	}
	if len(stack) > 0 {
		stack[len(stack)-1].Next = nil
	}
	return stack[0]
}
