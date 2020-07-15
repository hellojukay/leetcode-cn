#https://leetcode-cn.com/problems/find-pivot-index/
from typing import List;
class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        s = 0
        for n in nums:
            s = s + n
        tmp = 0
        for i,n in enumerate(nums):
            if  tmp*2 == (s-n):
                return i
            tmp = tmp + n
        return -1

