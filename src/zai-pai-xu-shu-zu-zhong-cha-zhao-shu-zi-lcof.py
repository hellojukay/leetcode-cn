#https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        dic = {}
        for i in nums:
            if dic.get(i,0) == 0:
                dic[i] = 1
            else:
                dic[i] = dic[i] + 1
        return dic.get(target,0)
