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

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var result *ListNode
	for _, list := range lists {
		result = mergeList(result, list)
	}
	return result
}

// merge two list
func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var tail *ListNode
	var head *ListNode
	var tmp *ListNode
	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 != nil && list2 != nil && list1.Val < list2.Val {
			tmp = list1
			list1 = list1.Next
			if tail == nil {
				tail = tmp
				head = tail
			} else {
				tail.Next = tmp
				tail = tail.Next
			}
			continue
		}
		if list1 != nil && list2 != nil && list1.Val >= list2.Val {
			tmp = list2
			list2 = list2.Next
			if tail == nil {
				tail = tmp
				head = tail
			} else {
				tail.Next = tmp
				tail = tail.Next
			}
			continue
		}
		if list1 == nil && list2 != nil {
			tmp = list2
			list2 = list2.Next
			tail.Next = tmp
			tail = tail.Next
			continue
		}
		if list1 != nil && list2 == nil {
			tmp = list1
			list1 = list1.Next
			tail.Next = tmp
			tail = tail.Next
			continue
		}
	}
	return head
}

func mergeKLists2(lists []*ListNode) *ListNode {

	var head *ListNode
	var tail *ListNode
	for len(lists) > 0 {
		index := findMin(lists)
		if index < 0 {
			break
		}
		if head == nil {
			head = lists[index]
			tail = head
		} else {
			tail.Next = lists[index]
			tail = tail.Next
		}
		lists[index] = lists[index].Next
		if lists[index] == nil {
			lists = append(lists[:index], lists[index+1:]...)
		}
	}
	return head
}

func findMin(lists []*ListNode) int {
	var index = -1
	for i := range lists {
		if lists[i] != nil {
			if index < 0 {
				index = i
				continue
			}
			if index >= 0 && lists[i].Val < lists[index].Val {
				index = i
			}
		}

	}
	return index
}
