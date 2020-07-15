#https://leetcode-cn.com/problems/palindrome-linked-list-lcci/
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        nums = []
        while head != None:
            nums.append(head.val)
            head = head.next
        for i in range(len(nums)):
            if nums[i] != nums[-1-i]:
                return False
            if i >= len(nums)-i-1:
                break
        return True

