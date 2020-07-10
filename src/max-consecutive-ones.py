# https://leetcode-cn.com/problems/max-consecutive-ones/
class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        result = 0
        tmp = 0
        for i in nums:
            if i == 0:
                if tmp > result:
                    result == tmp
                tmp = 0
                continue
            tmp = tmp + 1
            if tmp > result:
                result = tmp
        return result
