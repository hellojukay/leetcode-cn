class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        m = {}
        for x in nums:
            count = m.get(x,0)
            m[x] = count + 1
        return sum(filter(lambda x: m.get(x) ==1, nums))
