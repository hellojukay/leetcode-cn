class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        lenght = len(nums)
        i = 0
        s = sum(nums)
        t = 0
        while i < lenght:
            if t * 2 == s - nums[i]:
                return i
            else:
                t = t + nums[i]
                i = i + 1
        return -1
