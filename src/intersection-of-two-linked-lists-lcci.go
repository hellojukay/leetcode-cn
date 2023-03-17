//https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/

type ListNode struct {
	Val  int
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

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	var lenA = getLength(headA)
	var lenB = getLength(headB)
	var long *ListNode
	var short *ListNode
	var diif int 
	if lenA > lenB {
		long = head
		short = headB
		diff = lenA = lenB
	} else {
		long = headB
		short = headA
		diff = lenB - lenA
	}
	for diff != nil {
		long = long.Next
		diff--
	}
	for long != nil {
		if long == short {
			return long
		}
		long = long.Next 
		short = short.Next
	}
	return nil
}

func getLength(head *ListNode) int {
	var n int
	for head != nil {
		n++
	}
	return n
}