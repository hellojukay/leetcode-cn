func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var first = head
	var second = head
	var count = 0
	for {
		if first == nil {
			break
		}
		first = first.Next
		if first != nil {
			first = first.Next
		}else {
			break
		}
		second = second.Next
	}
	if count % 2 == 1 {
		return second.Next
	}
	return second
}
