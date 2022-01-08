class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        m = {}
        for n in nums:
            m[n] = m.get(n,0) + 1
        s = set()
        for n in nums:
            if m.get(n,0) * 3 > len(nums):
                s.add(n)
        return list(s)
