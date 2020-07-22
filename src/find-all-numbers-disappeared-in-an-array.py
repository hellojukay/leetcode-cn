class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        for n in nums:
            n = abs(n)
            if nums[n-1] < 0:
                continue
            else:
                nums[n-1] = nums[n-1] * -1
        result = []
        for i,n in enumerate(nums):
            if n > 0:
                result.append(i+1)
        return result

