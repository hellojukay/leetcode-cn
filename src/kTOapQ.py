# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class BSTIterator:

    def __init__(self, root: TreeNode):
        self.root= root
        def toList(root):
            result = []
            if not root:
                return []
            else:
                result.extend(toList(root.left))
                result.append(root.val)
                result.extend(toList(root.right))
            return result
        self.current = toList(root)


    def next(self) -> int:
        n = self.current[0]
        self.current = self.current[1:]
        return n


    def hasNext(self) -> bool:
        if len(self.current) > 0:
            return True
        return False




# Your BSTIterator object will be instantiated and called as such:
# obj = BSTIterator(root)
# param_1 = obj.next()
# param_2 = obj.hasNext()
