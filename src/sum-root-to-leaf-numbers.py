# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumNumbers(self, root: TreeNode) -> int:
        def foo(root: TreeNode, curent: int, result: List[int]) -> None:
            if root == None:
                result.append(curent)
                return
            if not root.left and not root.right:
                curent = curent * 10 + root.val
                result.append(curent)
                return
            if root.left:
                foo(root.left, curent * 10 + root.val, result)
            if root.right:
                foo(root.right, curent * 10 + root.val, result)

        result = []
        foo(root, 0, result)
        return sum(result)
