# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def lowestCommonAncestor(
        self, root: "TreeNode", p: "TreeNode", q: "TreeNode"
    ) -> "TreeNode":
        def include(root: "TreeNode", p: "TreeNode") -> bool:
            if root == p:
                return True
            if root == None:
                return False
            return include(root.left, p) or include(root.right, p)

        if include(root.left, p) and include(root.left, q):
            return self.lowestCommonAncestor(root.left, p, q)
        if include(root.right, p) and include(root.right, q):
            return self.lowestCommonAncestor(root.right, p, q)
        return root
