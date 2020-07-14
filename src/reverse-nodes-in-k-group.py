# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        if head == None:
            return None
        stack =[]
        queue = []
        p = head
        count = 0 # 节点总数
        while head != None:
            stack.append(head)
            if len(stack) == k:
                while len(stack) > 0:
                    queue.append(stack.pop())
            head = head.next
            count = count + 1
        while len(stack) > 0:
            queue.append(stack.pop(0))
        if count < k:
            return p
        p = queue[0]
        for i in range(len(queue)):
            if i == 0:
                continue
            p.next = queue[i]
            p = p.next
        p.next = None
        return queue[0]
