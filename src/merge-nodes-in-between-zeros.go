/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var dummy = &ListNode{Next: nil, Val: 0}
	var p = dummy
	var c = 0
	for {
		if head == nil {
			break
		}
		if head.Val == 0 || head.Next == nil {
			if c != 0 {
				p.Next = &ListNode{Next: nil, Val: c}
				p = p.Next
				c = 0
			}
		} else {
			c = c + head.Val
		}
		head = head.Next
	}
	return dummy.Next
}