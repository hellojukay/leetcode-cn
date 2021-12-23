from typing import List


class Solution:
    def findLengthOfLCIS(self, nums: List[int]) -> int:
        if len(nums) <= 1:
            return len(nums)
        length = len(nums)
        m = 1
        count = 1
        i = 0
        while i < length:
            if i + 1 <= length-1:
                if nums[i] < nums[i+1]:
                    count = count + 1
                    if count > m:
                        m = count
                else:
                    count = 1
                i = i + 1
            else:
                break
        return m


s = Solution()
print(s.findLengthOfLCIS([1, 3, 5, 4, 7]))
print(s.findLengthOfLCIS([1, 3, 5, 7]))
