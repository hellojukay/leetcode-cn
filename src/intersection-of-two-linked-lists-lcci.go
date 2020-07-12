//https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var m = make(map[*ListNode]bool)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if m[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}