# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def findSecondMinimumValue(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """

        def toList(root, result):
            if not root:
                return result
            result.append(root.val)
            result = toList(root.left, result)
            result = toList(root.right, result)
            return result

        result = toList(root, [])
        if len(result) < 2:
            return -1
        result = sorted(result)
        n = result[0]
        for x in result:
            if x > n:
                return x
        return -1
