class Solution:
    def thirdMax(self, nums: List[int]) -> int:
        m = max(nums)
        s1 = []
        for n in nums:
            if n < m :
                s1.append(n)
        if len(s1) == 0:
            return m
        second = max(s1)
        s2 = []
        for n in s1:
            if n < second:
                s2.append(n)
        if len(s2) == 0:
            return m
        return max(s2)
