# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def search(self, root: TreeNode, list: List[int], k: int) -> List[int]:
        if len(list) > k:
            return list
        if root.right:
            list = self.search(root.right, list, k)
        list.append(root.val)
        if len(list) > k:
            return list
        if root.left:
            list = self.search(root.left, list, k)

        return list

    def kthLargest(self, root: TreeNode, k: int) -> int:
        list = self.search(root, [], k)
        return list[k-1]
