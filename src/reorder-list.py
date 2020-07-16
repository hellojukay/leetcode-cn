# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: ListNode) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        stack = []
        while head:
            stack.append(head)
            head = head.next
        p = None
        tail = None
        while len(stack):
            if not p:
                p = stack.pop(0)
                tail = p
            if len(stack):
                tail.next = stack.pop(-1)
                tail = tail.next
            if len(stack):
                tail.next = stack.pop(0)
                tail = tail.next
        if tail:
            tail.next = None
        return p
