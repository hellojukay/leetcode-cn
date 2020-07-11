#https://leetcode-cn.com/problems/find-majority-element-lcci/
from typing import List;
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        dic = {}
        m = 0
        number = -1
        for n in nums:
            count = dic.get(n,0)
            if count == 0:
                count = 1
            else:
                count = count + 1
            dic[n] = count
            if count > m:
               m = count
               number = n
        print(m,number)
        if m  * 2 > len(nums):
            return number
        else:
            return -1
s = Solution()
s.majorityElement([3,2,3])

