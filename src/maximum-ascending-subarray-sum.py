class Solution:
    def maxAscendingSum(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        i = 0
        m = nums[0]
        t = m
        while i + 1 < len(nums):
            if nums[i] < nums[i + 1]:
                t = t + nums[i + 1]
                if t > m:
                    m = t
            else:
                t = nums[i + 1]
                if t > m:
                    m = t
            i = i + 1
        return m
