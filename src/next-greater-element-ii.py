#https://leetcode-cn.com/problems/next-greater-element-ii/
class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        if len(nums) == 1:
            return [-1]
        result = []
        for i , n in enumerate(nums):
            if i == len(nums)-1:
                i = -1
            tmp = -1
            for offset in range(1,len(nums)):
                j = (i + offset) % len(nums)
                if nums[j] > n:
                    tmp = nums[j]
                    break
            result.append(tmp)
        return result
