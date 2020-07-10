class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        dic = {}
        result = []
        for n in nums1:
            count = dic.get(n,0)
            if count == 0:
                dic[n] = 1
            else:
                dic[n] = count+1
        for n in nums2:
            count = dic.get(n,0)
            if count > 0 :
                result.append(n)
                dic[n] = count-1
        return result