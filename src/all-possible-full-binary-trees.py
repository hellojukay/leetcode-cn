# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def allPossibleFBT(self, n: int) -> List[TreeNode]:
        if n % 2 == 0:
            return []
        if n == 1:
            return [TreeNode(0, None, None)]
        result = []
        i = 1
        while i < n and n - 1 - i > 0:
            # 左子树是一棵 i 个节点的满二叉树
            leftList = self.allPossibleFBT(i)
            # 有子树是一颗 n - 1 -i 节点的满二叉树
            rightList = self.allPossibleFBT(n-1-i)
            for left in leftList:
                for right in rightList:
                    result.append(TreeNode(0, left, right))
        return result
