package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	if head == nil {
		return 0
	}
	var length int
	var p = head
	for p != nil {
		length++
		p = p.Next
	}
	var stack []int = make([]int, 0, length/2)
	i := 0
	p = head
	for i < length/2 {
		i++
		stack = append(stack, p.Val)
		p = p.Next
	}

	var n = length/2 - 1
	var max = p.Val + stack[n]
	for n >= 0 {
		tmp := p.Val + stack[n]
		if tmp > max {
			max = tmp
		}
		p = p.Next
		n--
	}
	return max
}
