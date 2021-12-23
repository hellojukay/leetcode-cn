class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        m = {}
        for n in nums:
            x = m.get(n, 0)
            m[n] = x+1
            if x+1 > int(len(nums)/2):
                return n
        return 0
