# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        stack = []
        c = 0
        while l1 or l2:
            s = 0
            if l1:
                s = s + l1.val 
                l1 = l1.next
            if l2:
                s = s + l2.val
                l2 = l2.next
            s = s + c
            c = int(s/10)
            stack.append(s%10)
        if c != 0:
            stack.append(c)
        head = None
        p = None
        for n in stack:
            node = ListNode(n)
            if head == None:
                head = node
                p = head
            else:
                p.next = node
                p = p.next
        return head

