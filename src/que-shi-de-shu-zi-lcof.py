class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        for n in range(len(nums)):
            if n != nums[n]:
                return n
        return nums[-1]+1

