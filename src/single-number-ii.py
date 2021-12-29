class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        m = {}
        for n in nums:
            count = m.get(n, 0)
            m[n] = count + 1
        for n in nums:
            count = m.get(n, 0)
            if count == 1:
                return n
        return 0
