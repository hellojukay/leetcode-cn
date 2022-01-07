/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var s1 []*ListNode
    var s2 []*ListNode
    for l1 != nil{
        s1 = append(s1,l1)
        l1 = l1.Next
    }
    for l2 != nil{
        s2 = append(s2,l2)
        l2 = l2.Next
    }
    var c = 0 
    var head *ListNode
    for len(s1) > 0 && len(s2) > 0 {
        var n = s1[len(s1)-1].Val +s2[len(s2)-1].Val + c
        c = n / 10
        var node = &ListNode{
            Val:n%10,
        }
        if head == nil{
            head = node
        } else {
            node.Next = head
            head = node
        }
        s1 = s1[:len(s1)-1]
        s2 = s2[:len(s2)-1]
    }
    for len(s1) > 0 {
        var n = s1[len(s1)-1].Val + c
        c = n / 10
        var node = &ListNode{
            Val:n%10,
        }
        if head == nil{
            head = node
        } else {
            node.Next = head
            head = node
        }
        s1 = s1[:len(s1)-1]
    }
    for len(s2) > 0 {
        var n = s2[len(s2)-1].Val + c
        c = n / 10
        var node = &ListNode{
            Val:n%10,
        }
        if head == nil{
            head = node
        } else {
            node.Next = head
            head = node
        }
        s2 = s2[:len(s2)-1]
    }
    if c > 0 {
        var node = &ListNode{
            Val:c,
        }
        if head == nil{
            head = node
        } else {
            node.Next = head
            head = node
        }
    }
    return head
}

