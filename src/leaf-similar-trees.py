class Solution(object):
    def leafSimilar(self, root1, root2):
        """
        :type root1: TreeNode
        :type root2: TreeNode
        :rtype: bool
        """

        def toList(root, result):
            if not root:
                return result
            if (not root.right) and (not root.left):
                result.append(root.val)
                return result
            if root.left:
                result = toList(root.left, result)
            if root.right:
                result = toList(root.right, result)
            return result

        def check(nums1, nums2):
            if len(nums1) != len(nums2):
                return False
            i = 0
            while i < len(nums1):
                if nums1[i] != nums2[i]:
                    return False
                i = i + 1
            return True

        nums1 = toList(root1, [])
        nums2 = toList(root2, [])
        return check(nums1, nums2)
