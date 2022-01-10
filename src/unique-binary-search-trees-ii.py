# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def generateTrees(self, n: int) -> List[TreeNode]:
        def dfs(start: int,end: int) -> List[TreeNode]:
            if start > end:
                return [None]
            i = start
            result = []
            while i <= end:
                lefts = dfs(start,i-1)
                rights = dfs(i+1,end)
                for left in lefts:
                    for right in rights:
                        root = TreeNode(i)
                        root.left = left
                        root.right = right
                        result.append(root)
                i = i + 1
            return result


        reslult = dfs(1,n)
        return reslult
