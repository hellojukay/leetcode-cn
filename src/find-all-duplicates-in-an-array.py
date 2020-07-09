class Solution:
    def findDuplicates(self, nums: List[int]) -> List[int]:
        dic = {}
        result = []
        for n in nums:
            if dic.get(n,False):
                result.append(n)
                continue
            dic[n] = True
        return result
