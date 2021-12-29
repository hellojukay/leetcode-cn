# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.left = left
class Solution:
    def insert(self, root: TreeNode, result: TreeNode) -> TreeNode:
        if root == None:
            return result
        if root and (not root.right) and (not root.left):
            root.right = result
            return root
        if root.right:
            result = self.insert(root.right, result)
        if root.left:
            result = self.insert(root.left, result)
        root.left = None
        root.right = result
        return root

    def flatten(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if root == None:
            return
        list = self.insert(root, None)
        if list:
            root.left = None
            root.right = list.right
