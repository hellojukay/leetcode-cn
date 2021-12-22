# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
# https://leetcode-cn.com/problems/aMhZSa/submissions/
class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        if head == None:
            return True
        s = ''
        while True:
            if head == None:
                break
            s += str(head.val)
            head = head.next
        return s == s[::-1]
