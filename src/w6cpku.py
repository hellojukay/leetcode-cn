# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class Solution:
    def toList(self, root: TreeNode, result: List[int]) -> list[int]:
        if not root:
            return result
        result.append(root.val)
        result = self.toList(root.left, result)
        return self.toList(root.right, result)

    def getSum(self, n: int, nums: List[int]) -> int:
        count = 0
        for m in nums:
            if m >= n:
                count = count + m
        return count

    def setNode(self, root: TreeNode, nums: List[int]) -> None:
        if not root:
            return
        root.val = self.getSum(root.val, nums)
        self.setNode(root.left, nums)
        self.setNode(root.right, nums)

    def convertBST(self, root: TreeNode) -> TreeNode:
        list = self.toList(root, [])
        if not root:
            return root
        self.setNode(root, list)
        return root
