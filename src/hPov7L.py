# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def largestValues(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        curent = [root]
        next = []
        result = []
        while len(curent):
            tmp = []
            for node in curent:
                tmp.append(node.val)
                if node.left:
                    next.append(node.left)
                if node.right:
                    next.append(node.right)
            curent = next
            next = []
            m = max(tmp)
            result.append(m)
        return result
