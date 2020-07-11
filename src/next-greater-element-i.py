#https://leetcode-cn.com/problems/next-greater-element-i/
class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        result = []
        for index in range(len(nums1)):
            tmp = -1
            flag = False
            for j in range(len(nums2)):
                if nums1[index] == nums2[j]:
                    flag = True
                    continue
                if flag and nums2[j] > nums1[index]:
                    tmp = nums2[j]
                    break
            result.append(tmp)
        return result

