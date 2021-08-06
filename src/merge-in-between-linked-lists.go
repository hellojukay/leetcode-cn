package main
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	// find list2 tail
	list2_tail := list2
	for {
		if list2_tail.Next == nil {
			break
		}
		list2_tail = list2_tail.Next
	}
	// find b point
	end := list1
	for i := 0; i <= b; i++ {
		end = end.Next
	}
	list2_tail.Next = end

	// find start
	start := list1
	for i := 0; i < (a-1); i++ {
		start = start.Next
	}
	start.Next = list2
	return list1
}
