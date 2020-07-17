class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        dic = {}
        for n in nums:
            count = dic.get(n,0)
            dic[n] = count + 1
        for n in nums:
            count = dic.get(n)
            if count == 1:
                return n
        return 0
