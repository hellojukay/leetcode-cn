# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        queue = []
        p = head
        for i in range(1,m):
            queue.append(p)
            p = p.next
        stack = []
        for i in range(m,n+1):
            stack.append(p)
            p = p.next
        while len(stack):
            queue.append(stack.pop(-1))
        while p:
            queue.append(p)
            p = p.next
        head = None
        p = None
        for node in queue:
            if not head:
                head = node
                p = head
                continue
            else:
                p.next = node
                p = p.next
        if p:
            p.next = None
        return head
