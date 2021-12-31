from typing import List


class Solution:
    def numOfPairs(self, nums: List[str], target: str) -> int:
        count = 0
        i = 0
        lenght = len(nums)
        while i < lenght:
            j = 0
            while j < lenght:
                if i == j:
                    j = j + 1
                    continue
                if (nums[i] + nums[j]) == target:
                    count = count + 1
                j = j + 1
            i = i + 1
        return count


s = Solution()
print(s.numOfPairs(["777", "7", "77", "77"], "7777"))
