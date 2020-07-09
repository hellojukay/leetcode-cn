class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        result = {}
        dic = {}
        for n in nums1:
            dic[n] = True
        for n in nums2:
            if dic.get(n,False):
                result[n] = True
        return result.keys()