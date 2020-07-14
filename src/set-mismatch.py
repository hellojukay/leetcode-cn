#https://leetcode-cn.com/problems/set-mismatch/
class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        dic = {}
        result = []
        s = int((1+len(nums)) * len(nums) / 2)
        tmp = 0
        for n in nums:
            count = dic.get(n,0)
            if count == 1:
                result.append(n)
            count = count+1
            dic[n]= count
            tmp  = tmp + n
        tmp = tmp - result[0]
        tmp = s - tmp
        result.append(tmp)
        return result
