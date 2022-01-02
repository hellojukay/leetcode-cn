import re


class Solution:
    def areNumbersAscending(self, s: str) -> bool:
        nums = re.findall(r"\d+", s)
        i = 1
        while i < len(nums):
            if int(nums[i]) <= int(nums[i - 1]):
                return False
        return True
