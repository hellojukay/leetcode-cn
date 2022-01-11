class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        m = {}
        for n in nums:
            m[n] = m.get(n,0) + 1
        for key in m.keys():
            if m.get(key,0) == 1:
                return key
        return 0
