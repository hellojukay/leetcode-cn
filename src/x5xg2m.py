# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverse(self, head: ListNode, result: ListNode):
        if head == None:
            return result
        next = head.next
        head.next = result
        return self.reverse(next, head)

    def reverseList(self, head: ListNode) -> ListNode:
        if head == None:
            return head
        return self.reverse(head, None)
