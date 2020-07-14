#https://leetcode-cn.com/problems/number-of-good-pairs/
class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        count = 0
        dic = {}
        for i in range(len(nums)):
            slice = dic.get(nums[i],[])
            slice.append(i)
            dic[nums[i]] = slice
        for i in range(len(nums)):
            slice = dic.get(nums[i],[])
            for n in slice:
                if i < n:
                    count = count + 1
        return count
