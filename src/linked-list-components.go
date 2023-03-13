/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
	var result int
	var count int
	var hash = make(map[int]bool)
	for _, n := range nums {
		hash[n] = true
	}
	for head != nil {
		if _, ok := hash[head.Val]; ok {
			count++
			head = head.Next
			continue
		}
		if count > 0 {
			result++
		}
		count = 0
		head = head.Next
	}
	if count > 0 {
		result++
	}
	return result
}