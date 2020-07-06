
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        j = 0
        for index in range(len(nums)):
            if nums[index] != 0 :
                continue
            if j <= index :
                j = index + 1
            # 找到不为 0 的数来补充
            while j < len(nums):
                if nums[j] != 0:
                    tmp = nums[j]
                    nums[j] = 0
                    nums[index] = tmp
                    break
                j = j + 1

