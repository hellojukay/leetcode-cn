/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	var p = head
	var result []int
	for {
		if p == nil {
			break
		}
		var q = p.Next
		for {
			if q == nil {
				result = append(result, 0)
				break
			}
			if q.Val > p.Val {
				result = append(result, q.Val)
				break
			} else {
				q = q.Next
			}
		}
		p = p.Next
	}
	return result
}

