#https://leetcode-cn.com/problems/check-permutation-lcci/
class Solution:
    def CheckPermutation(self, s1: str, s2: str) -> bool:
        dic1 = {}
        dic2 = {}
        for ch in s1:
            count = dic1.get(ch,0)
            count = count+1
            dic1[ch] =count
        for ch in s2:
            count = dic2.get(ch,0)
            count = count+1
            dic2[ch] =count
        if len(s1) != len(s2):
            return False
        for ch,count in dic1.items():
            if dic2.get(ch,0) != count:
                return False
        return True
