# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def toSet(self, root: TreeNode, s: set) -> set:
        if root == None:
            return s
        s = self.toSet(root.left, s)
        s = self.toSet(root.right, s)
        s.add(root.val)
        return s

    def findTarget(self, root: TreeNode, k: int) -> bool:
        s = set()
        s = self.toSet(root, s)
        for n in s:
            if k-n != n and (k-n) in s:
                return True
        return False
