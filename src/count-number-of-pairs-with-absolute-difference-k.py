class Solution:
    def countKDifference(self, nums: List[int], k: int) -> int:
        length = len(nums)
        i = 0
        count = 0
        while i < length:
            j = i + 1
            while j < length:
                if abs(nums[i]-nums[j]) == k:
                    count = count + 1
                j = j + 1
            i = i + 1
        return count
