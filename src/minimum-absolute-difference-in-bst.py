# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def getMinimumDifference(self, root: TreeNode) -> int:
        def toList(root: TreeNode) -> List[int]:
            if not root:
                return []
            result = []
            result.extend(toList(root.left))
            result.append(root.val)
            result.extend(toList(root.right))
            return result
        nums = toList(root)
        length = len(nums)
        i = 0 
        result = []
        while i < length and i + 1 < length:
            result.append(abs(nums[i]-nums[i+1]))
            i = i + 1
        return min(result)
