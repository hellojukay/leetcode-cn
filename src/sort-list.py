# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
from typing import Optional


class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        def mergeSort(head: Optional[ListNode]) -> Optional[ListNode]:
            if not head:
                return head
            if not head.next:
                return head
            if not head.next.next:
                tmp = head.next
                head.next = None
                return merge(head,tmp)
            first , second = head,head
            while first:
                first = first.next
                if first:
                    first = first.next
                else:
                    break
                second = second.next
            tmp = second.next
            second.next = None
            first = mergeSort(head)
            second = mergeSort(tmp)
            return merge(first,second)
        def merge(root1: Optional[ListNode],root2: Optional[ListNode]) -> Optional[ListNode]:
            if not root1:
                return root2
            if not root2:
                return root1
            head = None
            if root1.val < root2.val:
                head = root1
                root1 = root1.next
            else:
                head = root2
                root2 = root2.next
            p = head
            while root1 and root2:
                if root1.val < root2.val:
                    p.next = root1
                    root1 = root1.next
                else:
                    p.next = root2
                    root2 = root2.next
                p = p.next
            while root1:
                p.next = root1
                p = p.next
                root1 = root1.next
            while root2:
                p.next = root2
                p = p.next
                root2 = root2.next
            p.next = None
            return head
        return mergeSort(head)


