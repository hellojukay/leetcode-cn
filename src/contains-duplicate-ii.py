from typing import List


class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        m = {}
        length = len(nums)
        i = 0
        while i < length:
            index = m.get(nums[i], None)
            if index != None and nums[i] == nums[index] and abs(i-index) <= k:
                return True
            else:
                m[nums[i]] = i
            i = i + 1
        return False


s = Solution()
print(s.containsNearbyDuplicate([1, 2, 3, 1], 3))
